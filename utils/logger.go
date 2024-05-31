package utils

import (
	"sync"

	"github.com/threatwinds/logger"
)

var (
	agentLogger        *logger.Logger
	loggerOnceInstance sync.Once
)

// CreateLogger returns a single instance of a Logger configured to save logs to a rotating file.
func CreateLogger(filename string) *logger.Logger {
	loggerOnceInstance.Do(func() {
		agentLogger = logger.NewLogger(
			&logger.Config{Format: "text", Level: 100, Output: filename, Retries: 3, Wait: 5},
		)
	})
	return agentLogger
}
