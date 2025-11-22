package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Entry -> Struct representasi dari log
func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "eko").Info("Hello Logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "khannedy")
	entry.Info("Hello Entry")
}
