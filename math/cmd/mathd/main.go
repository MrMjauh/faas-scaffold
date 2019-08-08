package main

import (
	"flag"
	"github.com/MrMjauh/faas-scaffold/commons/pkg/mux"
	"github.com/MrMjauh/faas-scaffold/math/internal/pkg/handler"
	"log"
	"net/http"
)

func main(){
	var port = flag.String("port", "8080", "sets the port to serve requests from")
	var mode = flag.String("mode", "multiply", "supports mode math or add")
	flag.Parse()

	log.Println("Starting up math service on port " + *port)

	router := mux_common.CreateRoutingTemplate()
	apiVersion1Routes := mux_common.CreateAPIRoute(router, "v1")
	if *mode == "multiply" {
		apiVersion1Routes.Handle("/multiply", mux_common.WrappedHandler(http.HandlerFunc(handler.MultiplyHandler))).Methods("GET")
	} else if *mode == "add" {
		apiVersion1Routes.Handle("/add", mux_common.WrappedHandler(http.HandlerFunc(handler.AdditionHandler))).Methods("GET")
	} else {
		log.Fatal("Invalid mode")
	}
	apiVersion1Routes.Handle("/panic", mux_common.WrappedHandler(http.HandlerFunc(handler.PanicHandle))).Methods("GET")

	http.ListenAndServe(":" + *port, router)
}