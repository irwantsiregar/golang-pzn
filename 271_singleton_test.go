package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)


func TestSingletong(t *testing.T) {
	logrus.Info("Hello World");
	logrus.Error("Hello World");

	logrus.SetFormatter(&logrus.JSONFormatter{});

	logrus.Info("Hello World");
	logrus.Error("Hello World");
}
	
// As default gogrus Function is a initialized from a variable global 'std' 