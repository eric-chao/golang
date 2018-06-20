package db

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"adhoc/adhoc_uploader/config"
)

var (
	Pool *redis.Pool
)

func init() {
	host := config.GlobalConfig.Redis.Host
	port := config.GlobalConfig.Redis.Port
	redisHost := fmt.Sprintf("%s:%s", host, port)
	Pool = newPool(redisHost)
	close()
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 300 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Get(key string) (string, error) {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func HGet(key string, field string) (string, error) {
	conn := Pool.Get()
	defer conn.Close()
	data, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		return data, fmt.Errorf("error hget key %s, field %s: %v", key, field, err)
	}
	return data, err
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
