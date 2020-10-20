package main

import (
	"zap-learn/logger"
)

func main() {
	if err := logger.Setup(); err != nil {
		logger.Error(err)
		return
	}
	logger.FDebug("prefix", "key1", "value1", "key2", "value2")
	logger.FInfo("Base setup", "config path")
}
