package main

import (
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/consumer"
	. "adhoc/adhoc_data_fast/logger"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	fmt.Println(GlobalConfig.Kafka.Bootstrap)
	fmt.Println(GlobalConfig.Kafka.Topic)
	fmt.Println(GlobalConfig.Redis.Host)
	fmt.Println(GlobalConfig.Redis.Port)

	Logger.Info("begin...")
	Consume()
	Logger.Info("end...")
}

func goroutineLimitTest() {
	//runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	c := make(chan bool, 100)
	t := time.Tick(time.Second)

	go func() {
		for {
			select {
			case <-t:
				watching()
			}
		}
	}()

	for i := 0; i < 10000000; i++ {
		c <- true
		go worker(i, c)
	}

	fmt.Println("Done")
}

func watching() {
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
}

func worker(i int, c chan bool) {
	//fmt.Println("worker", i)
	time.Sleep(100 * time.Microsecond)
	<-c
}
