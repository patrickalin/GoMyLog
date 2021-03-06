package GoMylog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var Trace *log.Logger
var Info *log.Logger
var Warning *log.Logger
var Error *log.Logger

type Level int

const (
	ERROR Level = 1 + iota
	WARNING
	INFO
	TRACE
)

func Init(logLevel Level) {

	var (
		errorHandle   io.Writer
		infoHandle    io.Writer
		warningHandle io.Writer
		traceHandle   io.Writer
	)

	switch logLevel {
	case ERROR:
		errorHandle = os.Stdout
		warningHandle = ioutil.Discard
		infoHandle = ioutil.Discard
		traceHandle = ioutil.Discard
	case WARNING:
		errorHandle = os.Stdout
		warningHandle = os.Stdout
		infoHandle = ioutil.Discard
		traceHandle = ioutil.Discard
		log.Println("Log Level Warning")
	case INFO:
		errorHandle = os.Stdout
		warningHandle = os.Stdout
		infoHandle = os.Stdout
		traceHandle = ioutil.Discard
		log.Println("Log Level Info")
	case TRACE:
		errorHandle = os.Stdout
		warningHandle = os.Stdout
		infoHandle = os.Stdout
		traceHandle = os.Stdout
		log.Println("Log Level Trace")
	default:
		log.Fatal("Error in Log Level")
	}
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
