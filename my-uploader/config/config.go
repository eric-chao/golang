package config

import (
	"log"
	"gopkg.in/ini.v1"
	"os"
	"adhoc/adhoc_uploader/model"
)

var GlobalConfig model.Config

func init() {
	GlobalConfig, _ = NewConfig("config.ini")
	address := os.Getenv("ADDRESS")
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	if address != "" {
		GlobalConfig.Storage.Address = address
	}
	if host != "" {
		GlobalConfig.Redis.Host = host
	}
	if port != "" {
		GlobalConfig.Redis.Port = port
	}
}

func NewConfig(path string) (model.Config, error) {
	var config model.Config
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

