package redis

import (
	. "golang/my-kafka-redis/config"
	. "golang/my-kafka-redis/utils"
	"github.com/go-redis/redis"
	"fmt"
	"bytes"
	"os"
	"log"
)

var expRedis *redis.Client

func init() {
	redisAddress := fmt.Sprintf("%s:%s", GlobalConfig.Redis.ExpHost, GlobalConfig.Redis.ExpPort)
	expRedis = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: GlobalConfig.Redis.ExpPoolSize,
	})

	pong, err := expRedis.Ping().Result()
	if err != nil {
		log.Println("[redis-exp] \t", err)
		os.Exit(1)
	}

	log.Println("[redis-exp] \t", pong, "connected")
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
	value := expRedis.Get(buf.String()).Val()
	if value == "" {
		return "CONTROL"
	} else {
		return value
	}
}

func GetExpId(appId, clientId string, date int64) string {
	key := EncodeKey("adhoc_flaglog", "", UnixToTimeString(date, DefaultFormat), appId, clientId)
	value := expRedis.Get(key).Val()
	if value == "" {
		return "CONTROL"
	} else {
		return value
	}
}

func CloseExpRedis() {
	expRedis.Close()
}
