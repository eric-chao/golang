package process

import (
	. "golang/my-kafka-redis/logger"
	. "golang/my-kafka-redis/redis"
	. "golang/my-kafka-redis/utils"
	. "golang/my-kafka-redis/model"
	"time"
)

func APIStatLogProcess(body LogBody) {
	dataRedis := GetDataRedis()
	pipeline := dataRedis.Pipeline()

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

	// commit pipeline operation
	_, err := pipeline.Exec()
	if err != nil {
		Logger.Error("[redis] ", err)
	}
}
