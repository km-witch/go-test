package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// 0 Console & LogFile
// 1 LogFile only
const logMode = 1

var myLogger *log.Logger

func init() {
	now := time.Now().Format("2006_01_02")
	fpLog, err := os.OpenFile(now+".txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	if logMode == 0 {
		myLogger = log.New(io.MultiWriter(fpLog, os.Stdout), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		fmt.Println("log initialized in Console & LogFile Mode")
	} else {
		myLogger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		fmt.Println("log initialized in LogFile Only Mode ")
	}
}

func Log(msg string) {
	myLogger.Println(msg)
}
