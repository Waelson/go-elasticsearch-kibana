package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logFile, err := os.OpenFile("/app/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	logrus.SetOutput(logFile)
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Starting application...")

	// Simulando logs sendo enviados a cada 10 segundos
	for {
		logrus.Info("This is an info log message.")
		time.Sleep(10 * time.Second)
	}
}
