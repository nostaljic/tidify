package devlog

import (
	"log"
	"sync"
)

type devLogger struct {
	mu    sync.Mutex
	level string
}

var logger devLogger = devLogger{
	level: "Release",
}

func SetLogLevel(level string) {
	logger.mu.Lock()
	defer logger.mu.Unlock()
	logger.level = level
}

func Debug(msg ...interface{}) {
	if logger.level != "Develop" {
		return
	}

	log.Println(msg...)
}

func Info(msg ...interface{}) {
	log.Println(msg...)
}

func Fatal(msg ...interface{}) {
	log.Fatalln(msg...)
}

func Panic(msg ...interface{}) {
	log.Panicln(msg)
}
