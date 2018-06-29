package config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Adhoc struct {
		IgnoreDays int `ini:"ignore-days"`
		ExpireDays int `ini:"expire-days"`
	}

	Kafka struct {
		Bootstrap string `ini:"bootstrap.servers"`
		GroupId   string `ini:"group.id"`
		Topic     string `ini:"topic"`
	}

	Redis struct {
		ExpHost      string `ini:"exp.host"`
		ExpPort      string `ini:"exp.port"`
		ExpPoolSize  int    `ini:"exp.poolSize"`
		DataHost     string `ini:"data.host"`
		DataPort     string `ini:"data.port"`
		DataPoolSize int    `ini:"data.poolSize"`
	}

	MongoDB struct {
		Host string `ini:"host"`
		Port string `ini:"port"`
		DB   string `ini:"db"`
	}

	Log struct {
		Path string `ini:"path"`
	}

	Go struct {
		MaxProcs            int `ini:"max.procs"`
		MaxProcessGoroutine int `ini:"max.process.goroutine"`
	}
}

var GlobalConfig Config

func init() {
	GlobalConfig, _ = NewConfig("config.ini")
	kafkaBootstrap := os.Getenv("KAFKA_BOOTSTRAP")
	kafkaGroupId := os.Getenv("	KAFKA_GROUP_ID")
	expRedisHost := os.Getenv("EXP_REDIS_HOST")
	expRedisPort := os.Getenv("EXP_REDIS_PORT")
	expRedisPoolSize := os.Getenv("EXP_REDIS_POOL_SIZE")
	dataRedisHost := os.Getenv("DATA_REDIS_HOST")
	dataRedisPort := os.Getenv("DATA_REDIS_PORT")
	dataRedisPoolSize := os.Getenv("DATA_REDIS_POOL_SIZE")
	if kafkaBootstrap != "" {
		GlobalConfig.Kafka.Bootstrap = kafkaBootstrap
	}
	if kafkaGroupId != "" {
		GlobalConfig.Kafka.GroupId = kafkaGroupId
	}
	//ENV config for exp redis
	if expRedisHost != "" {
		GlobalConfig.Redis.ExpHost = expRedisHost
	}
	if expRedisPort != "" {
		GlobalConfig.Redis.ExpPort = expRedisPort
	}
	if expRedisPoolSize != "" {
		GlobalConfig.Redis.ExpPoolSize, _ = strconv.Atoi(expRedisPoolSize)
	}
	//ENV config for data redis
	if dataRedisHost != "" {
		GlobalConfig.Redis.DataHost = dataRedisHost
	}
	if dataRedisPort != "" {
		GlobalConfig.Redis.DataPort = dataRedisPort
	}
	if dataRedisPoolSize != "" {
		GlobalConfig.Redis.DataPoolSize, _ = strconv.Atoi(dataRedisPoolSize)
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
