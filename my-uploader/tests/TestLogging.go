package main

import (
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var Log log.Logger

func init() {
	Log := log.New()
	Log.SetLevel(log.InfoLevel)
	Log.Formatter = &log.TextFormatter{}

	// make sure GlobalConfig.Storage.Log ended with '/'
	Log.Out = &lumberjack.Logger{
		Filename:   fmt.Sprintf("%slogging-test.log", "/storage/logs/"),
		MaxSize:    30, 	// megabytes
		MaxBackups: 7,
		MaxAge:     1,     	// days
		Compress:   false, 	// disabled by default
		LocalTime:  true,
	}
}

func main() {
	
}