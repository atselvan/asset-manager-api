package main

import (
	"github.com/atselvan/go-utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	var a Asset
	utils.Logger{Message: "Starting the server..."}.Info()
	err := a.Init()
	if err != nil {
		utils.Logger{Message: appInitErrorStr}.Error()
	}
}

func main() {

	r := mux.NewRouter()
	r.PathPrefix(apiPathPrefix).Path(healthApiPath).HandlerFunc(healthHandler).Methods("GET")
	r.PathPrefix(apiPathPrefix).Path(assetCategoryApiPath).HandlerFunc(assetCategoryHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(assetTypeApiPath).HandlerFunc(assetTypeHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(assetBrandApiPath).HandlerFunc(assetBrandHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(assetStatusApiPath).HandlerFunc(assetStatusHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(assetsApiPath).HandlerFunc(assetsHandler).Methods("GET", "POST", "PUT")
	r.NotFoundHandler = http.HandlerFunc(pageNotFoundHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
