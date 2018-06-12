package consumer

import (
	. "adhoc/adhoc_data_fast/config"
	. "adhoc/adhoc_data_fast/logger"
	. "adhoc/adhoc_data_fast/model"
	. "adhoc/adhoc_data_fast/redis"
	. "adhoc/adhoc_data_fast/process"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strings"
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
			Logger.Debugf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			Logger.Debugf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

}

func parseBody(msg []byte) {
	p := &RequestBody{}
	json.Unmarshal(msg, p)

	appId := GetAppId(p.AppKey)
	if appId == "" {
		//do nothing
	} else {
		clientId := p.ClientId
		fromSystem := strings.HasPrefix(clientId, "ADHOC_EXP:")
		for _, stat := range p.Stats {
			var expIds []string
			if len(stat.ExperimentIds) > 0 {
				expIds = stat.ExperimentIds
			} else {
				date := stat.Timestamp
				if fromSystem {
					expIds = append(expIds, strings.Split(clientId, ":")[1])
				} else {
					expIds = append(expIds, GetExpId(appId, clientId, date))
				}
			}

			modId := GetModId(appId, clientId)
			for _, expId := range expIds {
				log := LogBody{
					Timestamp:  stat.Timestamp,
					AppId:      appId,
					ExpId:      expId,
					ModId:      modId,
					StatKey:    stat.Key,
					StatValue:  stat.Value,
					AccValue:   stat.AccValue,
					Summary:    p.Summary,
					Custom:     p.Custom,
					FromSystem: fromSystem,
				}
				//process log
				AllCounter.NewLogProcess(log)
			}
		}
	}

	<-channel
}
