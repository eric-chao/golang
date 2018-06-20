package process

import (
	. "adhoc/adhoc_data_fast_golang/redis"
	. "adhoc/adhoc_data_fast_golang/utils"
	. "adhoc/adhoc_data_fast_golang/model"
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
		redisClient := NewDataRedisClient()
		defer redisClient.Close()

		redisClient.HIncrBy(apiKey, EncodeKey("", "", expId), 1)
	}
}
