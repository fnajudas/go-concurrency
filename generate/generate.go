package generate

import (
	"fmt"
	"go-concurrency/strings"
	"go-concurrency/util/logger"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

const (
	totalFile     = 200
	contentLength = 300
)

var tempPath = filepath.Join(os.Getenv("TEMP"), "chapter-A.59-pipeline-temp")

func GenerateFiles() (err error) {
	logger := logger.New()
	loggers := logger.WithFields(logrus.Fields{
		"Layer":     "generate",
		"Func Name": "GenerateFiles",
	})
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < totalFile; i++ {
		fileName := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content, _ := strings.RandomString(contentLength)
		err := ioutil.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			loggers.Errorf("Error write file name in Generate Files")
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}

	log.Printf("%d of total files created", totalFile)
	return nil
}
