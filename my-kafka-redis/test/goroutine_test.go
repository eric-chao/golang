package test

import (
	"testing"
	"time"
	"runtime"
)

func Test_Goroutine_Limit(test *testing.T) {
	test.Log()
	//runtime.GOMAXPROCS(GlobalConfig.Go.MaxProcs)
	c := make(chan bool, 100)
	t := time.Tick(time.Second)

	go func() {
		for {
			select {
			case <-t:
				watching(test)
			}
		}
	}()

	for i := 0; i < 10000000; i++ {
		c <- true
		go worker(i, c)
	}

	test.Log("Done")
}

func watching(test *testing.T) {
	test.Logf("NumGoroutine: %d\n", runtime.NumGoroutine())
}

func worker(i int, c chan bool) {
	//fmt.Println("worker", i)
	time.Sleep(100 * time.Microsecond)
	<-c
}
