package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	. "adhoc/adhoc_data_fast_golang/config"
	. "adhoc/adhoc_data_fast_golang/logger"
	. "adhoc/adhoc_data_fast_golang/consumer"
	. "adhoc/adhoc_data_fast_golang/semaphore"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	print()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	//	Set log.connection.close=false to suppress these connection messages
	//	@see https://github.com/edenhill/librdkafka/issues/1089
	//	"auto.offset.reset": "earliest", default value is "latest"
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               GlobalConfig.Kafka.Bootstrap,
		"group.id":                        GlobalConfig.Kafka.GroupId,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"log.connection.close":            false,
	})

	if err != nil {
		Logger.Error("[kafka]: connect error!")
		return
	}
	defer c.Close()

	// control the consumer parallelism
	// channel := make(chan bool, GlobalConfig.Go.Goroutine)
	s := NewSemaphore(GlobalConfig.Go.Goroutine)

	topic := GlobalConfig.Kafka.Topic
	c.SubscribeTopics([]string{topic}, nil)
	run := true

	Logger.Info("process begin...")
	for run == true {
		select {
		case sig := <-sigchan:
			Logger.Infof("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				Logger.Errorf("%% %v\n", e)
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				Logger.Errorf("%% %v\n", e)
				c.Unassign()
			case *kafka.Message:
				s.Acquire()
				go ParseBody(e.Value, s)
			case kafka.PartitionEOF:
				Logger.Infof("%% Reached %v\n", e)
			case kafka.Error:
				Logger.Errorf("%% Error: %v\n", e)
				run = false
			default:
				Logger.Infof("%% Ignored %v\n", e)
			}
		}
	}

	// wait for all goroutines completed.
	s.Wait()

	Logger.Info("process end...")

}

func print() {
	fmt.Println("[kafka-broker]\t", GlobalConfig.Kafka.Bootstrap)
	fmt.Println("[kafka-topic]\t", GlobalConfig.Kafka.Topic)
	fmt.Println("[redis-exp]\t", GlobalConfig.Redis.ExpHost)
	fmt.Println("[redis-exp]\t", GlobalConfig.Redis.ExpPort)
	fmt.Println("[redis-data]\t", GlobalConfig.Redis.DataHost)
	fmt.Println("[redis-data]\t", GlobalConfig.Redis.DataPort)
}
