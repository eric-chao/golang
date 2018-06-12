package process

import (
	. "adhoc/adhoc_data_fast/model"
	. "adhoc/adhoc_data_fast/utils"
	. "adhoc/adhoc_data_fast/redis"
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
		if now > logTime {
			logTime = now
		}

		if int(CalcAbs(now-logTime)/(3600*24)) > newLog.IgnoreDays {
			return
		}
	}

	//data redis client
	dataRedis := NewDataRedisClient()
	defer dataRedis.Close()
	pipeline := dataRedis.Pipeline()
	defer pipeline.Close()

	for _, customKey := range newLog.CustomKeys(body) {
		prefix := newLog.Prefix
		timeString := newLog.TimeString(logTime)

		resultField := EncodeKey("", "", body.StatKey, customKey)

		_, existed := commonStatKey[body.StatKey]
		if !existed || prefix == "adhoc_all" {
			if newLog.CountStat {
				sumKey := EncodeKey(prefix, "_sum", timeString, body.AppId)
				resultUvKey := EncodeKey(prefix, "_result_uv", timeString, body.AppId, body.StatKey, customKey)
				pipeline.HIncrByFloat(sumKey, resultField, body.StatValue)
				pipeline.PFAdd(resultUvKey, body.ClientId)
				pipeline.Exec()
			}
		}

		if !existed {
			sumSquareKey := EncodeKey(prefix, "_sum_square", timeString, body.AppId)
			if newLog.CountVariance && body.AccValue != 0.0 {
				newValue := body.AccValue
				oldValue := newValue - body.StatValue
				incSquareSum := (newValue * newValue) - (oldValue * oldValue)
				dataRedis.HIncrByFloat(sumSquareKey, resultField, incSquareSum)
			} else if newLog.CountVariance {
				historyKey := EncodeKey(prefix, "_history", timeString, body.AppId, body.ClientId, body.StatKey, customKey)
				newValue := dataRedis.IncrByFloat(historyKey, body.StatValue).Val()
				oldValue := newValue - body.StatValue
				incSquareSum := (newValue * newValue) - (oldValue * oldValue)
				// 1(s) = 1000000000(ns)
				pipeline.Expire(historyKey, time.Duration(newLog.ExpireDays*3600*24*1000000000))
				pipeline.HIncrByFloat(sumSquareKey, resultField, incSquareSum)
				pipeline.Exec()
			}
		}

		clientField := EncodeKey("", "", customKey)
		clientKey := EncodeKey(prefix, "_client", timeString, body.AppId)
		if newLog.UseLogLog {
			viewLogLogKey := EncodeKey(prefix, "_view_loglog", timeString, body.AppId, customKey)
			dataRedis.PFAdd(viewLogLogKey, body.ClientId)
			count := dataRedis.PFCount(viewLogLogKey).Val()
			dataRedis.HSet(clientKey, clientField, string(count))
		} else {
			viewKey := EncodeKey(prefix, "_view", timeString, body.AppId, body.ClientId, customKey)
			views := dataRedis.Incr(viewKey).Val()
			// 1(s) = 1000000000(ns)
			dataRedis.Expire(viewKey, time.Duration(newLog.ExpireDays*3600*24*1000000000))
			if views == 1 {
				dataRedis.HIncrByFloat(clientKey, clientField, 1)
			}
		}

	}

}
