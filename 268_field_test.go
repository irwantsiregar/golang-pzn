package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "khannedy").Info("Hello World")

	logger.WithField("username", "eko").
		WithField("name", "Eko Kurniawan").
		Info("Hello World")
}

// Adding Fields to in Log
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "eko",
		"name":     "Eko Kurniawan",
	}).Info("Hello World")
}
