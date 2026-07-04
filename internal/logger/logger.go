package logger

import (
	"log"
	"os"
)

var (
	Info  = log.New(os.Stdout, "[INFO]  ", log.LstdFlags)
	Error = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	Debug = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
)

func EnableDebug(enable bool) {
	if !enable {
		Debug.SetOutput(os.Stdout)
	}
}
