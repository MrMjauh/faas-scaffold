package main

import (
	"faas-scaffold/commons/pkg/mux_common"
	"faas-scaffold/multiply/internal/app/multiplyd/handler"
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "sets the port to serve requests from")

func setupV1Route(r* mux.Router) {
	r.Handle("/multiply", mux_common.WrappedHandler(http.HandlerFunc(handler.MultiplyHandler))).Methods("GET")
	r.Handle("/add", mux_common.WrappedHandler(http.HandlerFunc(handler.AdditionHandler))).Methods("GET")
}

func main(){
	log.Println("Starting up multiply service...")

	router := mux.NewRouter()
	// Common configuration make sures it looks the same on all micro-services
	mux_common.CreateRouteConfiguration(router)

	var apiVersion1Routes = mux_common.CreateRouteForVersion(router, "v1")
	setupV1Route(apiVersion1Routes)

	http.ListenAndServe(":" + *port, router)
}