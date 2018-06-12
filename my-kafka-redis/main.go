package main

import (
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/consumer"
	. "adhoc/adhoc_data_fast/logger"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	fmt.Println(GlobalConfig.Kafka.Bootstrap)
	fmt.Println(GlobalConfig.Kafka.Topic)
	fmt.Println(GlobalConfig.Redis.ExpHost)
	fmt.Println(GlobalConfig.Redis.ExpPort)

	Logger.Info("begin...")
	Consume()
	Logger.Info("end...")
}

