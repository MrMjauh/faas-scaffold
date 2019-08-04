package handler

import (
	"faas-scaffold/commons/pkg/rest"
	"faas-scaffold/faas-gateway/internal/pkg/dto"
	"faas-scaffold/faas-gateway/internal/pkg/service"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type RouteHandler struct {
	RoutesMutex *sync.RWMutex
	Routes * dto.ServiceRoutes
	ProxyResolver * service.ProxyResolver
}

const (
	ERROR_CODE_INVALID_ROUTE = 10000
)

func (routerHandler* RouteHandler) ProxyHandler(w http.ResponseWriter, r *http.Request) {
	routerHandler.RoutesMutex.RLock()
	defer routerHandler.RoutesMutex.RUnlock()

	host, urlRewrite, err := (*routerHandler.ProxyResolver).ResolveProxy(r.URL.String(), &routerHandler.Routes.Services)

	if err != nil {
		rest_common.WriteJsonError(w, rest_common.GeneralErrorResponse(ERROR_CODE_INVALID_ROUTE, err.Error()))
		return
	}

	// Rewrite the url where we strip of the multiply
	r.URL, _ = url.Parse(urlRewrite)
	u, _ := url.Parse("http://" + host)
	httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
}

func (routerHandler* RouteHandler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	routerHandler.RoutesMutex.RLock()
	defer routerHandler.RoutesMutex.RUnlock()
	rest_common.WriteJsonResponse(w, routerHandler.Routes)
}