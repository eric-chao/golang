package process

import (
	. "adhoc/adhoc_data_fast_golang/logger"
	. "adhoc/adhoc_data_fast_golang/model"
	. "adhoc/adhoc_data_fast_golang/utils"
	. "adhoc/adhoc_data_fast_golang/redis"
	"time"
)

type NewLog struct {
	Prefix        string
	CountStat     bool
	CountVariance bool
	UseLogLog     bool
	IgnoreDays    int
	ExpireDays    int
	CustomNames   []string
	CustomKeys    func(body LogBody) []string
	TimeString    func(t int64) string
}

var commonStatKey = map[string]string{
	"Event-session":                   "",
	"Event-GET_EXPERIMENT_FLAGS_HTTP": "",
	"Event-GET_EXPERIMENT_FLAGS":      "",
	"Event-duration":                  "",
	"Event-GET_SDK_CONFIG":            "",
}

func (newLog *NewLog) NewLogProcess(body LogBody) {
	logTime := body.Timestamp
	now := time.Now().Unix()
	if !body.FromSystem {
		if logTime > now {
			logTime = now
		}

		if int(CalcAbs(now-logTime)/(3600*24)) > newLog.IgnoreDays {
			return
		}
	}

	//data redis client
	dataRedis := DataRedisClient
	defer dataRedis.Close()
	pipeline := dataRedis.Pipeline()
	//defer pipeline.Close()

	for _, customKey := range newLog.CustomKeys(body) {
		prefix := newLog.Prefix
		timeString := newLog.TimeString(logTime)

		resultField := EncodeKey("", "", body.StatKey, customKey)
		if newLog.CountStat {
			sumKey := EncodeKey(prefix, "_sum", timeString, body.AppId)
			resultUvKey := EncodeKey(prefix, "_result_uv", timeString, body.AppId, body.StatKey, customKey)
			pipeline.HIncrByFloat(sumKey, resultField, body.StatValue)
			pipeline.PFAdd(resultUvKey, body.ClientId)
			_, err := pipeline.Exec()
			if err != nil {
				Logger.Error("[redis] ", err)
			}
		}
		var accValue float64
		//Logger.Infof("[process] %s, %s, %f, %t", prefix, body.StatKey, body.AccValue, body.AccValue == accValue)
		sumSquareKey := EncodeKey(prefix, "_sum_square", timeString, body.AppId)
		if newLog.CountVariance && body.AccValue != accValue {
			newValue := body.AccValue
			oldValue := newValue - body.StatValue
			incSquareSum := (newValue * newValue) - (oldValue * oldValue)
			dataRedis.HIncrByFloat(sumSquareKey, resultField, incSquareSum)
		} else if newLog.CountVariance {
			historyKey := EncodeKey(prefix, "_history", timeString, body.AppId, body.ClientId, body.StatKey, customKey)
			newValue := pipeline.IncrByFloat(historyKey, body.StatValue)
			pipeline.Expire(historyKey, time.Duration(newLog.ExpireDays*3600*24*1000000000))
			_, err := pipeline.Exec()
			if err != nil {
				Logger.Error("[redis] ", err)
			}
			oldValue := newValue.Val() - body.StatValue
			incSquareSum := (newValue.Val() * newValue.Val()) - (oldValue * oldValue)
			// 1(s) = 1000000000(ns)
			dataRedis.HIncrByFloat(sumSquareKey, resultField, incSquareSum)
		}

		clientField := EncodeKey("", "", customKey)
		clientKey := EncodeKey(prefix, "_client", timeString, body.AppId)
		if newLog.UseLogLog {
			viewLogLogKey := EncodeKey(prefix, "_view_loglog", timeString, body.AppId, customKey)
			pipeline.PFAdd(viewLogLogKey, body.ClientId)
			count := pipeline.PFCount(viewLogLogKey)
			_, err := pipeline.Exec()
			if err != nil {
				Logger.Error("[redis] ", err)
			}
			dataRedis.HSet(clientKey, clientField, count.Val())
		} else {
			viewKey := EncodeKey(prefix, "_view", timeString, body.AppId, body.ClientId, customKey)
			views := pipeline.Incr(viewKey)
			// 1(s) = 1000000000(ns)
			pipeline.Expire(viewKey, time.Duration(newLog.ExpireDays*3600*24*1000000000))
			_, err := pipeline.Exec()
			if err != nil {
				Logger.Error("[redis] ", err)
			}
			if views.Val() == 1 {
				dataRedis.HIncrByFloat(clientKey, clientField, 1)
			}
		}

	}

}
