package consumer

import (
	. "adhoc/adhoc_data_fast_golang/logger"
	. "adhoc/adhoc_data_fast_golang/model"
	. "adhoc/adhoc_data_fast_golang/redis"
	. "adhoc/adhoc_data_fast_golang/process"
	. "adhoc/adhoc_data_fast_golang/semaphore"
	"encoding/json"
	"strings"
	"time"
	"sync"
)

func ParseBody(msg []byte, s *Semaphore) {
	defer s.Release()
	startTime := time.Now()

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
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					defer wg.Done()
					AllCounter.NewLogProcess(log)
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					HourlyCounter.NewLogProcess(log)
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					DailyCounter.NewLogProcess(log)
				}()

				wg.Add(1)
				go func() {
					defer wg.Done()
					MonthlyUvCounter.NewLogProcess(log)
				}()
				//StatCounter.NewLogProcess(log)
				//ApiCounter.NewLogProcess(log)
				wg.Add(1)
				go func() {
					defer wg.Done()
					APIStatLogProcess(log)
				}()

				wg.Wait()
			}
		}
	}

	Logger.Infof("[redis] takes: %d nanoseconds", time.Now().Sub(startTime))
	// <-channel
	// s.Release()
}
