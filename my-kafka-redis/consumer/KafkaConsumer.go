package consumer

import (
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/logger"
	. "adhoc/adhoc_data_fast/model"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
)

//to control the parallelism
var channel = make(chan bool, GlobalConfig.Go.Goroutine)

func Consume() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": GlobalConfig.Kafka.Bootstrap,
		"group.id":          GlobalConfig.Kafka.GroupId,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		Logger.Error("[kafka]: connect error!")
		return
	}

	topic := GlobalConfig.Kafka.Topic
	c.SubscribeTopics([]string{topic}, nil)
	defer c.Close()

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			channel <- true
			go parseBody(msg.Value)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

}

func parseBody(msg []byte) {
	p := &RequestBody{}
	json.Unmarshal(msg, p)
	Logger.Infof("%s - %s", p.AppKey, p.ClientId)
	<- channel
}
