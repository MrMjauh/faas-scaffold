package service

import (
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"
	"sync"
)

func UpdateRegistry(configurationChan <- chan dto.ServiceRoutes, routesMutex *sync.RWMutex, routes * dto.ServiceRoutes) {
	// If this loop ends badly, always make sure we unlock write lock
	defer routesMutex.Unlock()
	for true {
		data := <-configurationChan
		routesMutex.Lock()
		routes.ProviderName = data.ProviderName
		routes.Services = data.Services
		routesMutex.Unlock()
	}
}
