package process

import (
	. "golang/my-kafka-redis/redis"
	. "golang/my-kafka-redis/utils"
	. "golang/my-kafka-redis/model"
	"time"
)

type Api struct {
	NewLog
}

var ApiCounter = &Api{
	NewLog{
		Prefix:      "adhoc_daily_api",
		TimeString: func(t int64) string {
			return Daily(t)
		},
	},
}

func (api Api) NewLogProcess(body LogBody) {
	expId := body.ExpId
	if expId != "CONTROL" {
		now := time.Now().Unix()
		logTime := body.Timestamp
		if logTime > now {
			logTime = now
		}

		apiKey := EncodeKey(api.Prefix, "_count", api.TimeString(logTime), body.AppId)
		dataRedis := GetDataRedis()

		dataRedis.HIncrBy(apiKey, EncodeKey("", "", expId), 1)
	}
}
