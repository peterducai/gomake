package main

import (
	"os"
	"path"
	"sync"
	"time"
)

//check check error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	INFO = iota
	DEBUG
	WARNING
	ERROR
)

type logmessage struct {
	title, message string
	logtype        int
}

type logconfiguration struct {
	logfile string
	logdir  string
}

var mu sync.Mutex

//MicroLogger logger itself
type MicroLogger interface {
	Configure(conf *logconfiguration)
	Error(message string)
	Warn(message string)
	Deprecated(message string)
	Info(message string)
	Debug(message string)
}


//Error error message to log
func Error(message string) {
	msg := new(logmessage)
	WriteToFile(message)
}

//Warn warning message to log
func Warn(message string) {
	WriteToFile(message)
}

//Info info message to log
func Info(message string) {
	WriteToFile(message)
}

//Debug debug message to log
func Debug(message string) {
	WriteToFile(message)
}

//Deprecated deprecated message to log
func Deprecated(message string) {
	WriteToFile(message)
}

//getTime get current time
func getTime() string {
	var timeFormat = "2006/01/02 - 15:04:05"
	now := time.Now().Format(timeFormat)
	return now
}

/AppendPath get absolute path
func AppendPath(name string) string {
	wd, _ := os.Getwd()
	return path.Join(wd, name)
}

//WriteToFile write to file
func WriteToFile(text string) {
	mu.Lock()
	defer mu.Unlock()

	//LogFile := AppendPath(LogFile)
	//f, err := os.OpenFile(LogFile, os.O_APPEND|os.O_WRONLY, 0644)
	//defer f.Close()

	//t, err := f.WriteString(text)
	//check(err)
	//fmt.Printf("wrote %d bytes\n", t)

	//f.Sync()
}