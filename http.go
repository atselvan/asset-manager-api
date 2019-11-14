package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

// TODO: Add documentation in code

const (
	successStatusCode             = 200
	createdStatusCode             = 201
	badRequestStatusCode          = 400
	unauthorizedStatusCode        = 401
	notFoundStatusCode            = 404
	methodNotAllowedStatusCode    = 405
	internalServerErrorStatusCode = 500

	successStatus             = "200 OK"
	createdStatus             = "201 Created"
	badReqStatus              = "400 Bad Request"
	unauthorizedStatus        = "401 Unauthorized"
	notFoundStatus            = "404 Not Found"
	methodNotAllowed          = "405 Method Not Allowed"
	internalServerErrorStatus = "500 Internal Server Error"

	jsonMarshalErrorStr   = "JSON Marshal Error"
	jsonUnmarshalErrorStr = "JSON Unmarshal Error"
	apiAuthErrorStr       = "401 unauthorized. Please pass username and password to the API"
)

type Response struct {
	Message string `json:"message"`
}

type ErrResponse struct {
	Error string `json:"error"`
}

func getRequesterIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func basicAuthCheck(r *http.Request) error {
	var err error
	if r.Header.Get("Authorization") == "" {
		err = errors.New(apiAuthErrorStr)
	}
	return err
}

func ReadRequestBody(r *http.Request, out interface{}) error {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(reqBody) > 0 {
		err := json.Unmarshal(reqBody, out)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeHTTPResp(r *http.Request, w http.ResponseWriter, responseCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		writeErrorResp(r, w, internalServerErrorStatusCode, err)
	}
	w.WriteHeader(responseCode)
	_, err = w.Write(out)
	if err != nil {
		writeErrorResp(r, w, internalServerErrorStatusCode, err)
	}
}

func writeInfoResp(r *http.Request, w http.ResponseWriter, responseCode int, response string) {
	writeHTTPResp(r, w, responseCode, Response{response})
	Logger{r, response}.Info()
}

func writeWarnResp(r *http.Request, w http.ResponseWriter, successCode int, response string) {
	writeHTTPResp(r, w, successCode, Response{response})
	Logger{r, response}.Warn()
}

func writeErrorResp(r *http.Request, w http.ResponseWriter, s int, err error) {
	response := ErrResponse{err.Error()}
	writeHTTPResp(r, w, s, response)
	Logger{r, err.Error()}.Error()
}
