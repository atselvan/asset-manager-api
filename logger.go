package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	logFormat     = "| %s"
	httpLogFormat = "| %s | %s | %s | %s"
)

type Logger struct {
	Request *http.Request
	Message interface{}
}

func (l Logger) Info() {
	var out string
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	if l.Request != nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	infoLog.Println(out)
}

func (l Logger) Warn() {
	var out string
	warnLog := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	if l.Request == nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	warnLog.Println(out)
}

func (l Logger) Error() {
	var out string
	errLog := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	if l.Request == nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	errLog.Println(out)
}