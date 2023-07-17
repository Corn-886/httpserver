package util

import (
	"log"
)

var severityName = []string{
	0: "INFO",
	1: "DEBUG",
	2: "ERROR",
}

var logging loggingS

type loggingS struct {
	logLevel  int
	logparent string
}

func PrintDebugLog(msg string) {
	if logging.logLevel >= 1 {
		log.Println("DEBUG:" + logging.logparent + msg)
	}

}

func PrintINFOLog(format string, s ...string) {
	if logging.logLevel >= 0 {
		log.Println("INFO:"+logging.logparent+format, s)
	}

}

func PrintERRORLog(msg string) {
	if logging.logLevel >= 0 {
		log.Panicln("ERROR:" + logging.logparent + "")
	}

}

func init() {
	var envLoglevel = "INFO"
	for i, v := range severityName {
		if v == envLoglevel {
			logging.logLevel = i
		}
	}
	logging.logparent = "----------"
}
