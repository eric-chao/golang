package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	. "adhoc/adhoc_data_fast_golang/config"
)

var Logger *log.Logger

func init() {
	Logger = log.New()
	Logger.SetLevel(log.InfoLevel)
	Logger.Formatter = &log.TextFormatter{}

	// make sure GlobalConfig.Storage.Log ended with '/'
	Logger.Out = &lumberjack.Logger{
		Filename:   fmt.Sprintf("%sadhoc-data-fast.log", GlobalConfig.Log.Path),
		MaxSize:    30, 	// megabytes
		MaxBackups: 7,
		MaxAge:     1,     	//days
		Compress:   false, 	// disabled by default
		LocalTime:  true,
	}
}