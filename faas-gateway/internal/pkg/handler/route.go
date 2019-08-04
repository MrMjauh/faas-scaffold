package handler

import (
	"faas-scaffold/commons/pkg/rest"
	"faas-scaffold/faas-gateway/internal/pkg/dto"
	"faas-scaffold/faas-gateway/internal/pkg/service"
	"net/http"
	"sync"
)

type RouteHandler struct {
	RoutesMutex *sync.RWMutex
	Routes * dto.ServiceRoutes
	ProxyResolver * service.ProxyResolver
}

const (
	INVALID_ROUTE_CODE = 20000
)

func (routerHandler* RouteHandler) ProxyHandler(w http.ResponseWriter, r *http.Request) {
	routerHandler.RoutesMutex.RLock()
	defer routerHandler.RoutesMutex.RUnlock()

	redirectUrl, err := (*routerHandler.ProxyResolver).ResolveProxy(r.URL.Path, &routerHandler.Routes.Services)

	if err != nil {
		rest.WriteJsonError(w, rest.GeneralErrorResponse(INVALID_ROUTE_CODE, err.Error()))
		return
	}

	// Do redirect
	rest.WriteJsonResponse(w, redirectUrl)
}

func (routerHandler* RouteHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	rest.WriteJsonResponse(w, routerHandler.Routes)
}