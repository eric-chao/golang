package process

import (
	. "golang/my-kafka-redis/model"
	. "golang/my-kafka-redis/redis"
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
	dataRedis := GetDataRedis()
	dataRedis.SAdd(stat.Prefix+body.AppId, body.StatKey)
}
