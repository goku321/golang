package main

import (
	"fmt"
	"golang/exercises/logging/basic-logging/user"
	"os"
	"path"
	"runtime"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	logFile, _ := os.OpenFile("logfile", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(logFile)
	log.SetReportCaller(true)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("New logging")
	user.Name()
}
