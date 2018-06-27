package process

import (
	. "adhoc/adhoc_data_fast_golang/model"
	. "adhoc/adhoc_data_fast_golang/redis"
)

type Statistics struct {
	NewLog
}

var StatCounter = &Statistics{
	NewLog{
		Prefix: "adhoc_stat_",
	},
}

func (stat Statistics) NewLogProcess(body LogBody) {
	redisClient := DataRedisClient
	defer redisClient.Close()

	redisClient.SAdd(stat.Prefix + body.AppId, body.StatKey)

}
