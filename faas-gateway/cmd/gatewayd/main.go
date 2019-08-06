package main

import (
	"flag"
	"github.com/MrMjauh/faas-scaffold/docker/pkg"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/handler"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/provider"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/service"
	"log"
	"net/http"
	"sync"
	"time"
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

	go func() {
		dockerService.LinuxOnly_Me()
		time.Sleep(time.Second)
	}()

	allHandler := handler.RouteHandler{
		RoutesMutex: &routesMutex,
		Routes: &routes,
		ProxyResolver: &proxyResolver,
	}

	http.HandleFunc("/", allHandler.ProxyHandler)
	http.HandleFunc("/stats", allHandler.StatsHandler)

	http.ListenAndServe(":" + *port, nil)
}
