package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Create output log to become JSON formatter
func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.Warn("Hello Logging")
	logger.Error("Hello Logging")
}
