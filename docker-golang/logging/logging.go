package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var AppLog *LoggerClass

type LoggerClass struct {
	Log *logrus.Logger
}

func SetUpLogging() {
	AppLog = &LoggerClass{}
	AppLog.Log = logrus.New()
	/*file, err := os.OpenFile(
		"logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		AppLog.Log.Fatal(err)
	}*/

	AppLog.Log.SetOutput(os.Stdout)
	AppLog.Log.SetFormatter(&logrus.JSONFormatter{})

}

func (logger *LoggerClass) WriteLogWarn(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Warn(msg)
}

func (logger *LoggerClass) WriteLogError(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Error(msg)
}

func (logger *LoggerClass) WriteLogInfo(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Info(msg)
}

func (logger *LoggerClass) WriteLogDebug(msg string, fields map[string]interface{}) {
	logger.Log.WithFields(fields).Debug(msg)
}
