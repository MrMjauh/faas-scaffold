package service

import "github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"

type ProxyResolver interface {
	ResolveProxy(path string, services* map[string]dto.Service) (string, string, error)
}