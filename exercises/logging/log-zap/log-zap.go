package main

import (
	"golang/exercises/logging/log-zap/user"

	"go.uber.org/zap"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.OutputPaths = []string{
		"../basic-logging/logfile",
		"../basic-logging/logfilenew",
	}
	logger, _ := loggerConfig.Build()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
	logger.Info("failed to fetch URL")
	user.Name()
}
