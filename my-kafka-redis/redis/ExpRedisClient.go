package redis

import (
	. "adhoc/adhoc_data_fast_golang/config"
	. "adhoc/adhoc_data_fast_golang/utils"
	"github.com/go-redis/redis"
	"fmt"
	"bytes"
)

var ExpRedisClient *redis.Client

func init() {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.ExpHost, GlobalConfig.Redis.ExpPort)
	ExpRedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.ExpPoolSize,
	})
}

//get appId from appKey
func GetAppId(appKey string) string {
	if len(appKey) < 6 {
		return ""
	}
	return string([]byte(appKey)[6:])
}

func GetModId(appId, clientId string) string {
	var buf bytes.Buffer
	buf.WriteString("adhoc_mod_new_client_")
	buf.WriteString(appId)
	buf.WriteString("_")
	buf.WriteString(clientId)
	redisClient := ExpRedisClient
	defer redisClient.Close()

	return redisClient.Get(buf.String()).Val()
}

func GetExpId(appId, clientId string, date int64) string {
	redisClient := ExpRedisClient
	defer redisClient.Close()
	key := EncodeKey("adhoc_flaglog", "", UnixToTimeString(date, DefaultFormat), appId, clientId)
	value := redisClient.Get(key).Val()
	if value == "" {
		return "CONTROL"
	} else {
		return value
	}
}
