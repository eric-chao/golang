package main

import (
	. "adhoc/adhoc_data_fast_golang/config"
	. "adhoc/adhoc_data_fast_golang/consumer"
	. "adhoc/adhoc_data_fast_golang/logger"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	fmt.Println(GlobalConfig.Kafka.Bootstrap)
	fmt.Println(GlobalConfig.Kafka.Topic)
	fmt.Println("[redis-exp]", GlobalConfig.Redis.ExpHost)
	fmt.Println("[redis-exp-port]", GlobalConfig.Redis.ExpPort)
	fmt.Println("[redis-data]", GlobalConfig.Redis.DataHost)
	fmt.Println("[redis-data-port]", GlobalConfig.Redis.DataPort)

	Logger.Info("begin...")
	Consume()
	Logger.Info("end...")
}

