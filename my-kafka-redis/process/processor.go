package process

import (
	. "golang/my-kafka-redis/logger"
	. "golang/my-kafka-redis/model"
	. "golang/my-kafka-redis/redis"
	. "golang/my-kafka-redis/semaphore"
	"encoding/json"
	"strings"
	"time"
	"sync"
)

func ParseBody(msg []byte, s *Semaphore) {
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
			var wg sync.WaitGroup
			redisStartTime := time.Now()
			modId := GetModId(appId, clientId)
			for _, expId := range expIds {
				log := LogBody{
					Timestamp:  stat.Timestamp,
					AppId:      appId,
					ExpId:      expId,
					ModId:      modId,
					ClientId:   clientId,
					StatKey:    stat.Key,
					StatValue:  stat.Value,
					AccValue:   stat.AccValue,
					Summary:    p.Summary,
					Custom:     p.Custom,
					FromSystem: fromSystem,
				}
				Logger.Info("[LogBody] ", log.ToString())
				//process log
				wg.Add(1)
				go func(){
					AllCounter.NewLogProcess(log)
					wg.Done()
				}()

				wg.Add(1)
				go func(){
					HourlyCounter.NewLogProcess(log)
					wg.Done()
				}()

				wg.Add(1)
				go func(){
					DailyCounter.NewLogProcess(log)
					wg.Done()
				}()

				wg.Add(1)
				go func(){
					MonthlyUvCounter.NewLogProcess(log)
					wg.Done()
				}()

				wg.Add(1)
				go func(){
					APIStatLogProcess(log)
					wg.Done()
				}()
				//StatCounter.NewLogProcess(log)
				//ApiCounter.NewLogProcess(log)
				//APIStatLogProcess(log)
			}
			wg.Wait()
			redisElapsed := time.Since(redisStartTime)
			Logger.Infof("[redis] takes: %d ns", redisElapsed)
		}
	}

	//release
	s.Release()
}
