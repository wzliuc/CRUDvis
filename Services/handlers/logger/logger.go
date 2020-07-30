package logger

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var warnLogger *log.Logger
var errLogger *log.Logger

func init() {
	infoLogger = log.New(os.Stderr, "[INFO]    ", log.Ldate|log.Ltime)
	warnLogger = log.New(os.Stderr, "[WARNING] ", log.Ldate|log.Ltime)
	errLogger = log.New(os.Stderr, "[ERROR]   ", log.Ldate|log.Ltime)
}

// LogInfo logs information to console
func LogInfo(msg string) {
	infoLogger.Println(msg)
}

// LogWarn logs warning to console
func LogWarn(msg string) {
	warnLogger.Println(msg)
}

// LogErr logs error to console
func LogErr(msg interface{}) {
	errLogger.Println(msg)
}
