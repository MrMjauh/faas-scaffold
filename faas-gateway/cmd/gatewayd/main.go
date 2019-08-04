package main

import (
	docker "faas-scaffold/docker/pkg"
	"faas-scaffold/faas-gateway/internal/pkg/dto"
	"faas-scaffold/faas-gateway/internal/pkg/handler"
	"faas-scaffold/faas-gateway/internal/pkg/provider"
	"faas-scaffold/faas-gateway/internal/pkg/service"
	"flag"
	"log"
	"net/http"
	"sync"
)

var port = flag.String("port", "8080", "sets the port to serve requests from")

func main() {
	log.Println("Starting gateway")

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

	http.HandleFunc("/", allHandler.ProxyHandler)
	http.HandleFunc("/stats", allHandler.StatsHandler)

	http.ListenAndServe(":" + *port, nil)
}
