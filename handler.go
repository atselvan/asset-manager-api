package main

import (
	"fmt"
	"github.com/atselvan/go-pgdb-lib"
	"github.com/atselvan/go-utils"
	"net/http"
)

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
		utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, utils.NewError("enum name is a required parameter"))
		return
	}
	e := pgdb.Enum{
		Name: enumName,
	}
	err := e.Get()
	if err != nil {
		utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, err)
		return
	} else {
		if len(e.Values) < 1 {
			utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, utils.StringSlice{})
		} else {
			utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, e.Values)
		}
	}
}

func enumPostHandler(w http.ResponseWriter, r *http.Request, enumName string) {
	if enumName == "" {
		utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, utils.NewError("enum name is a required parameter"))
		return
	}
	value := r.URL.Query().Get("value")
	if value == "" {
		utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, utils.NewError("value is a required parameter"))
		return
	} else {
		e := pgdb.Enum{
			Name:   enumName,
			Values: utils.StringSlice{value},
		}
		err := e.Update()
		if err != nil {
			utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, err)
			return
		} else {
			utils.WriteInfoResp(w, r, utils.SuccessStatusCode, fmt.Sprintf("'%s' is added to %s", value, enumName))
		}
	}
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	var a Asset

	switch r.Method {

	case "GET":
		assets, err := a.Get()
		if err != nil {
			utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
		} else {
			fmt.Println(assets)
			fmt.Println(len(assets))
			if len(assets) < 1 {
				utils.WriteHTTPResp(w, r, utils.NotFoundStatusCode, utils.Response{Message: "No Assets were found"})
			} else {
				utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, &assets)
			}
		}

	case "POST":
		err := utils.ReadRequestBody(r, &a)
		if err != nil {
			utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
			return
		}
		if err := a.IsValid(); err != nil {
			utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, err)
			return
		}
		id, err := a.Exists()
		if id != "" {
			utils.WriteInfoResp(w, r, utils.FoundStatusCode, fmt.Sprintf("Asset with serial '%s' already exists with id '%s'", a.Serial, id))
		} else if id == "" {
			id, err := a.Add()
			if err != nil {
				utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
			} else {
				utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, utils.Response{Message: fmt.Sprintf("Asset information is added with id '%s'", id)})
			}
		} else {
			utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
		}
	}
}
