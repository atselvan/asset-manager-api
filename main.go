package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	var a Asset
	Logger{Message: "Starting the server..."}.Info()
	err := a.Init()
	if err != nil {
		Logger{Message: appInitErrorStr}.Error()
	}
}

// TODO : Move constants to constants file
const (
	apiPathPrefix   = "/api/v1"
	categoryApiPath = "/categories"
	typeApiPath     = "/types"
	brandApiPath    = "/brands"
	assetsApiPath   = "/assets"
)

func main() {

	r := mux.NewRouter()
	r.PathPrefix(apiPathPrefix).Path(categoryApiPath).HandlerFunc(categoryHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(typeApiPath).HandlerFunc(typeHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(brandApiPath).HandlerFunc(brandHandler).Methods("GET", "POST")
	r.PathPrefix(apiPathPrefix).Path(assetsApiPath).HandlerFunc(assetsHandler).Methods("GET", "POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
