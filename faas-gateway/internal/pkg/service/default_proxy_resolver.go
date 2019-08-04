package service

import (
	"errors"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"
	"regexp"
	"strconv"
)

type DefaultProxyResolver struct {

}

var proxyRegexMatcher = regexp.MustCompile("^/[a-z0-9]+")

func (resolver DefaultProxyResolver) ResolveProxy(path string, services *map[string]dto.Service) (string, string, error) {
	firstRouteUrl := proxyRegexMatcher.FindString(path)
	if firstRouteUrl == "" {
		return "", "", errors.New("invalid route url")
	}

	if firstRouteUrl == "/" {
		return "", "", errors.New("need a route key to route")
	}

	routeKey := firstRouteUrl[1:]
	service, serviceFound := (*services)[routeKey]

	if !serviceFound {
		return "", "", errors.New("no service found for key = " + routeKey)
	}

	return service.Alias + ":" + strconv.FormatUint(uint64(service.Port), 10), path[len(firstRouteUrl):], nil
}
