package main

import (
	"flag"
	"github.com/MrMjauh/faas-scaffold/commons/pkg/mux"
	"github.com/MrMjauh/faas-scaffold/docker/pkg"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/handler"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/provider"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/service"
	"log"
	"net/http"
	"sync"
)

func main() {
	var port = flag.String("port", "8080", "sets the port to serve requests from")
	flag.Parse()

	log.Println("Starting gateway on port " + *port)

	// Create our registry of routes
	var routesMutex sync.RWMutex
	var routes dto.ServiceRoutes
	// Create all services we need
	var dockerService docker.Docker = docker.DefaultImpl{}
	var proxyResolver service.ProxyResolver = service.DefaultProxyResolver{}

	provider := provider.DockerProvider{PollingIntervalMillis: 1000, DockerService: &dockerService}
	providerChan := make(chan dto.ServiceRoutes)

	// Run the providing function and listening function
	go provider.Provide(providerChan)
	go service.UpdateRegistry(providerChan, &routesMutex, &routes)

	allHandler := handler.RouteHandler{
		RoutesMutex: &routesMutex,
		Routes: &routes,
		ProxyResolver: &proxyResolver,
	}

	router := mux_common.CreateRoutingTemplate()
	apiVersion1Routes := mux_common.CreateAPIRoute(router, "v1")
	apiVersion1Routes.Handle("/stats", mux_common.WrappedHandler(http.HandlerFunc(allHandler.StatsHandler))).Methods("GET")
	router.PathPrefix("/").HandlerFunc(allHandler.ProxyHandler)
	http.ListenAndServe(":" + *port, router)
}
