package main

import (
	"faas-scaffold/commons/pkg/mux"
	"faas-scaffold/multiply/internal/pkg/handler"
	"flag"
	"log"
	"net/http"
)

func main(){
	var port = flag.String("port", "8080", "sets the port to serve requests from")
	flag.Parse()

	log.Println("Starting up multiply service on port " + *port)

	router := mux_common.CreateRoutingTemplate()
	apiVersion1Routes := mux_common.CreateAPIRoute(router, "v1")
	apiVersion1Routes.Handle("/multiply", mux_common.WrappedHandler(http.HandlerFunc(handler.MultiplyHandler))).Methods("GET")
	apiVersion1Routes.Handle("/add", mux_common.WrappedHandler(http.HandlerFunc(handler.AdditionHandler))).Methods("GET")

	http.ListenAndServe(":" + *port, router)
}