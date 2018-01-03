package main

import (
	"os"
	"sync"
	"time"
)

var mu sync.Mutex

//LogFile path to log file
var LogFile = "somepath"

//WriteToFile write to file
func WriteToFile(text string, f *os.File) {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	n, err := f.WriteString(text)
}

//Error error message to log
func Error() {

}

//Warn warning message to log
func Warn() {

}

//Info info message to log
func Info() {

}

//Debug debug message to log
func Debug() {

}

//Deprecated deprecated message to log
func Deprecated() {

}

//getTime get current time
func getTime() string {
	var timeFormat = "2006/01/02 - 15:04:05"
	now := time.Now().Format(timeFormat)
	return now
}
