package provider

import (
	docker "faas-scaffold/docker/pkg"
	"faas-scaffold/faas-gateway/internal/pkg/dto"
	"log"
	"strconv"
	"time"
)

const (
	PROVIDER_NAME = "docker"

	LABEL_PORT = "faas.port"
	LABEL_NAME = "faas.name"
)

type DockerProvider struct {
	PollingIntervalMillis uint64
	DockerService * docker.Docker
}

func (provider * DockerProvider) Provide(configurationChan chan <- dto.ServiceRoutes) {
	sleepFunc := func() {
		time.Sleep(time.Duration(provider.PollingIntervalMillis) * time.Millisecond)
	}

	for true {
		containers, err := (*provider.DockerService).ListAllContainers()
		if err != nil {
			log.Println(err)
			sleepFunc()
			continue
		}

		services := make(map[string]dto.Service)
		for _, container := range containers {
			if container.State != docker.STATUS_STATE_RUNNING {
				continue
			}

			portStr, portKeyFound := container.Labels[LABEL_PORT];
			name, nameKeyFound := container.Labels[LABEL_NAME];

			if !portKeyFound || !nameKeyFound {
				continue
			}

			port, err := strconv.ParseInt(portStr, 10, 16)
			if err != nil {
				continue
			}

			_, serviceFound := services[name]
			if serviceFound {
				continue
			}

			services[name] = dto.Service{
				Name: name,
				Port: uint16(port),
			}
		}

		routes := dto.ServiceRoutes{
			ProviderName: PROVIDER_NAME,
			Services: services,
		}
		configurationChan <- routes

		sleepFunc()
	}
}