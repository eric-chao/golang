package process

import (
	. "adhoc/adhoc_data_fast_golang/redis"
	. "adhoc/adhoc_data_fast_golang/utils"
	. "adhoc/adhoc_data_fast_golang/model"
	"time"
)

func APIStatLogProcess(body LogBody) {
	redisClient := NewDataRedisClient()
	defer redisClient.Close()
	pipeline := redisClient.Pipeline()
	defer pipeline.Close()

	expId := body.ExpId
	if expId != "CONTROL" {
		now := time.Now().Unix()
		logTime := body.Timestamp
		if logTime > now {
			logTime = now
		}

		apiKey := EncodeKey("adhoc_daily_api", "_count", func(t int64) string { return Daily(t) }(logTime), body.AppId)
		pipeline.HIncrBy(apiKey, EncodeKey("", "", expId), 1)
	}

	pipeline.SAdd("adhoc_stat_"+body.AppId, body.StatKey)
	pipeline.Exec()
}
