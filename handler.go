package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	//var c Category

	// TODO : Redesign HTTP response and Logger utils

	switch r.Method {

	case "GET":
		e := Enum{
			name: categoryEnumTypeName,
		}
		err := e.Get()
		if err != nil {
			writeErrorResp(r, w, badRequestStatusCode, err)
			return
		} else {
			writeHTTPResp(r, w, successStatusCode, e.values)
		}

	case "POST":
		value := r.URL.Query().Get("value")
		if value == "" {
			writeErrorResp(r, w, badRequestStatusCode, NewError("value is a required parameter"))
			return
		} else {
			e := Enum{
				name:   categoryEnumTypeName,
				values: StringSlice{value},
			}
			err := e.Update()
			if err != nil {
				writeErrorResp(r, w, badRequestStatusCode, err)
				return
			} else {
				writeInfoResp(r, w, successStatusCode, fmt.Sprintf("Catergory %s is added", value))
			}
		}
	}
}
