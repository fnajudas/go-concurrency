package start

import (
	"log"
	"time"

	"go-concurrency/generate"
	"go-concurrency/util/logger"

	"github.com/sirupsen/logrus"
)

func Start() {
	logger := logger.New()
	loggers := logger.WithFields(logrus.Fields{
		"Layer":     "start",
		"Func Name": "Start",
	})
	start := time.Now()

	err := generate.GenerateFiles()
	if err != nil {
		loggers.Errorf("Error getGenerateFiles in main.go")
		return
	}

	duration := time.Since(start)
	log.Println("Done in", duration.Seconds(), "seconds")
}
