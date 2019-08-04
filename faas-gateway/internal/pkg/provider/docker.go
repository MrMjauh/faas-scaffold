package provider

import (
	"github.com/MrMjauh/faas-scaffold/docker/pkg"
	"github.com/MrMjauh/faas-scaffold/faas-gateway/internal/pkg/dto"
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
		detailedMe, err := (*provider.DockerService).LinuxOnly_Me()
		if err != nil {
			log.Println(err)
			sleepFunc()
			continue
		}

		containers, err := (*provider.DockerService).GetContainers()
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

			alias := findAliasThatCanBeCalled(&detailedMe, &container)
			if alias == "" {
				continue
			}

			services[name] = dto.Service{
				Name: name,
				Port: uint16(port),
				Alias: alias,
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

func findAliasThatCanBeCalled(containerCallee * docker.DetailedContainer, containerToBeCalled * docker.Container) string {
	for calleeNetworkName, _ := range containerCallee.NetworkSettings.Networks {
		for toBeCalledNetworkName, _ := range containerCallee.NetworkSettings.Networks {
			if calleeNetworkName == toBeCalledNetworkName && len(containerCallee.NetworkSettings.Networks[toBeCalledNetworkName].Aliases) > 0 {
				// Just pick one, no need to know exactly
				return containerCallee.NetworkSettings.Networks[toBeCalledNetworkName].Aliases[0]
			}
		}
	}

	return ""
}