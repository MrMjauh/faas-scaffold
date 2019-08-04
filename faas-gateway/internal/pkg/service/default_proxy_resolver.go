package service

import (
	"errors"
	"faas-scaffold/faas-gateway/internal/pkg/dto"
	"regexp"
)

type DefaultProxyResolver struct {

}

var proxyRegexMatcher = regexp.MustCompile("^/[a-z0-9]+")

func (resolver DefaultProxyResolver) ResolveProxy(path string, services *map[string]dto.Service) (string, error) {
	firstRouteUrl := proxyRegexMatcher.FindString(path)
	if firstRouteUrl == "" {
		return "", errors.New("invalid route url")
	}

	if firstRouteUrl == "/" {
		return "", errors.New("need a route key to route")
	}

	routeKey := firstRouteUrl[1:]
	service, serviceFound := (*services)[routeKey]

	if !serviceFound {
		return "", errors.New("no service found for key = " + routeKey)
	}

	return service.Name, nil
}
