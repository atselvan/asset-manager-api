package main

import (
	"fmt"
	"net/http"
)

// TODO : Redesign HTTP response and Logger utils
// TODO: Added logging to init steps

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func categoryHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, categoryEnumTypeName)

	case "POST":
		enumPostHandler(w, r, categoryEnumTypeName)
	}
}

func typeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, typeEnumTypeName)

	case "POST":
		enumPostHandler(w, r, typeEnumTypeName)
	}
}

func brandHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, brandEnumTypeName)

	case "POST":
		enumPostHandler(w, r, brandEnumTypeName)
	}
}

func enumGetHandler(w http.ResponseWriter, r *http.Request, enumName string) {
	if enumName == "" {
		writeErrorResp(r, w, internalServerErrorStatusCode, NewError("enum name is a required parameter"))
		return
	}
	e := Enum{
		name: enumName,
	}
	err := e.Get()
	if err != nil {
		writeErrorResp(r, w, badRequestStatusCode, err)
		return
	} else {
		if len(e.values) < 1 {
			writeHTTPResp(r, w, successStatusCode, StringSlice{})
		} else {
			writeHTTPResp(r, w, successStatusCode, e.values)
		}
	}
}

func enumPostHandler(w http.ResponseWriter, r *http.Request, enumName string) {
	if enumName == "" {
		writeErrorResp(r, w, internalServerErrorStatusCode, NewError("enum name is a required parameter"))
		return
	}
	value := r.URL.Query().Get("value")
	if value == "" {
		writeErrorResp(r, w, badRequestStatusCode, NewError("value is a required parameter"))
		return
	} else {
		e := Enum{
			name:   enumName,
			values: StringSlice{value},
		}
		err := e.Update()
		if err != nil {
			writeErrorResp(r, w, badRequestStatusCode, err)
			return
		} else {
			writeInfoResp(r, w, successStatusCode, fmt.Sprintf("'%s' is added to %s", value, enumName))
		}
	}
}

func assetsHandeler(w http.ResponseWriter, r *http.Request) {

}
