package main

import (
	"fmt"
	"net/http"
)

// TODO : Redesign HTTP response and Logger utils
// TODO: Improve Handler info msg. eg: 'Apple1' is added to asset_brand
// TODO: Update home handler + add page not found handler

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

func assetsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		var a Asset
		assets, err := a.Get()
		if err != nil {
			writeErrorResp(r, w, internalServerErrorStatusCode, err)
		} else {
			fmt.Println(assets)
			fmt.Println(len(assets))
			if len(assets) < 1 {
				writeHTTPResp(r, w, notFoundStatusCode, Response{Message: "No Assets were found"})
			} else {
				writeHTTPResp(r, w, successStatusCode, &assets)
			}
		}

	case "POST":
		a := Asset{
			Name:     "MacBook Pro",
			Category: "device",
			Ctype:    "laptop",
			Model:    "MacBook Pro 15-inch SpaceGrey",
			Serial:   "C02VC1TBHTD51",
			Brand:    "Apple",
			MnfYear:  "2017",
			PDate:    "10/31/2017",
			Price:    "2799",
			Status:   "owned",
		}
		fmt.Println(a)

		id, err := a.Add()
		if err != nil {
			writeErrorResp(r, w, internalServerErrorStatusCode, err)
		} else {
			writeHTTPResp(r, w, successStatusCode, Response{Message: fmt.Sprintf("Asset information is added with id '%s'", id)})
		}
	}
}
