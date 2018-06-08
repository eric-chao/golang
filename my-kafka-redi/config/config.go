package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Redis struct {
		Host     string `ini:"host"`
		Port     string `ini:"port"`
		PoolSize int    `ini:"poolSize"`
	}

	Kafka struct {
		Bootstrap string `ini:"bootstrap.servers"`
		GroupId   string `ini:"group.id"`
		Topic     string `ini:"topic"`
	}

	Log struct {
		Path string `ini:"path"`
	}

	Go struct {
		MaxProcs  int `ini:"max.procs"`
		Goroutine int `ini:"goroutine"`
	}
}

var GlobalConfig Config

func init() {
	GlobalConfig, _ = NewConfig("config.ini")
	kafkaBootstrap := os.Getenv("KAFKA_BOOTSTRAP")
	kafkaGroupId := os.Getenv("	KAFKA_GROUP_ID")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPoolSize := os.Getenv("REDIS_POOL_SIZE")
	if kafkaBootstrap != "" {
		GlobalConfig.Kafka.Bootstrap = kafkaBootstrap
	}
	if kafkaGroupId != "" {
		GlobalConfig.Kafka.GroupId = kafkaGroupId
	}
	if redisHost != "" {
		GlobalConfig.Redis.Host = redisHost
	}
	if redisPort != "" {
		GlobalConfig.Redis.Port = redisPort
	}

	if redisPoolSize != "" {
		GlobalConfig.Redis.PoolSize, _ = strconv.Atoi(redisPoolSize)
	}
}

func NewConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path)
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	//解析成结构体
	conf.BlockMode = false
	err = conf.MapTo(&config)
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
