package logger

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	. "adhoc/adhoc_data_fast_golang/config"
)

var Logger *log.Logger

func init() {
	// format: %Y%m%d%H%M
	// make sure GlobalConfig.Storage.Log ended with '/'
	logFile := fmt.Sprintf("%sadhoc-data-fast.log.%s", GlobalConfig.Log.Path, "%Y%m%d")
	Logger = log.New()
	Logger.SetLevel(log.InfoLevel)
	Logger.Formatter = &log.TextFormatter{}
	// set Logger.Out to rotatelogs
	rl, _ := rotatelogs.New(logFile)
	Logger.Out = rl

}
