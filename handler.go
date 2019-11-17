package main

import (
	"fmt"
	"github.com/atselvan/go-pgdb-lib"
	"github.com/atselvan/go-utils"
	"net/http"
	"strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteInfoResp(w, r, utils.SuccessStatusCode, "OK")
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteInfoResp(w, r, utils.NotFoundStatusCode, fmt.Sprintf("Request path '%s' not found", r.RequestURI))
}

func assetCategoryHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, categoryEnumName)

	case "POST":
		enumPostHandler(w, r, categoryEnumName)
	}
}

func assetTypeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, typeEnumName)

	case "POST":
		enumPostHandler(w, r, typeEnumName)
	}
}

func assetBrandHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, brandEnumName)

	case "POST":
		enumPostHandler(w, r, brandEnumName)
	}
}

func assetStatusHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		enumGetHandler(w, r, statusEnumName)

	case "POST":
		enumPostHandler(w, r, statusEnumName)
	}
}

func enumGetHandler(w http.ResponseWriter, r *http.Request, enumName string) {
	e := pgdb.Enum{
		Name: "assets_" + enumName,
	}
	err := e.Get()
	if err != nil {
		utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, err)
		return
	} else {
		if len(e.Values) < 1 {
			utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, []string{})
			utils.Logger{Request: r, Message: fmt.Sprintf(getEnumSuccessStr, enumName)}.Info()
		} else {
			utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, e.Values)
			utils.Logger{Request: r, Message: fmt.Sprintf(getEnumSuccessStr, enumName)}.Info()
		}
	}
}

func enumPostHandler(w http.ResponseWriter, r *http.Request, enumName string) {
	value := r.URL.Query().Get("value")
	if value == "" {
		utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, utils.Error{ErrMsg: valueRequiredStr}.NewError())
		return
	} else {
		e := pgdb.Enum{
			Name: "assets_" + enumName,
		}
		if err := e.Get(); err != nil {
			utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
			return
		}
		if utils.EntryExists(e.Values, value) {
			utils.WriteInfoResp(w, r, utils.FoundStatusCode, fmt.Sprintf(enumValueExistsStr, enumName, value))
			return
		}
		e.Values = []string{value}
		err := e.Update()
		if err != nil {
			utils.WriteErrorResp(w, r, utils.BadRequestStatusCode, err)
			return
		} else {
			utils.WriteInfoResp(w, r, utils.SuccessStatusCode, fmt.Sprintf("%s '%s' is added", strings.Title(enumName), value))
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
			if len(assets) < 1 {
				utils.WriteHTTPResp(w, r, utils.NotFoundStatusCode, utils.Response{Message: noAssetFoundStr})
				utils.Logger{Request: r, Message: fmt.Sprintf(noAssetFoundStr)}.Info()
			} else {
				utils.WriteHTTPResp(w, r, utils.SuccessStatusCode, &assets)
				utils.Logger{Request: r, Message: fmt.Sprintf(getAssetSuccessStr)}.Info()
			}
		}

	case "POST":
		err := utils.ReadRequestBody(r, &a)
		if err != nil {
			utils.WriteErrorResp(w, r, utils.InternalServerErrorStatusCode, err)
			return
		}
		if err := a.IsValidAssetInfo(); err != nil {
			utils.WriteHTTPResp(w, r, utils.BadRequestStatusCode, err)
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
