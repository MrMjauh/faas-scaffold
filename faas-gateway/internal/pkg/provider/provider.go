package provider

import "faas-scaffold/faas-gateway/internal/pkg/dto"

// Inspired by https://github.com/containous/traefik/blob/master/pkg/provider/provider.go

type Provider interface {
	Provide(configurationChan chan<- dto.ServiceRoutes) error
	Init() error
}