package logging

import (
	"github.com/ICEBERG98/go-consul-cleanup/pkg/config"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func InitLogging() {
	setLoggingFormatter()
	setLogFile()
}
func setLogFile() {
	logFileName := config.Config.Logging.LogFile
	log.Infof("Using %s as Logfile", logFileName)
	logfile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panicf("Unable to open LogFile, Error- %s", err)
	}
	writer := io.Writer(logfile)
	if config.Config.Logging.LogToStdout == true {
		writer = io.MultiWriter(os.Stdout, logfile)
	}
	log.SetOutput(writer)
	log.SetReportCaller(true)
}
func setLoggingFormatter() {
	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
		// DisableLevelTruncation: true,
		PadLevelText: true,
		ForceQuote:   false,
	})
}
