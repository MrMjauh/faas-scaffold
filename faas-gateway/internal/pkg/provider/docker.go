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

		services, err := (*provider.DockerService).GetServices()
		if err != nil {
			log.Println(err)
			sleepFunc()
			continue
		}

		servicesDto := make(map[string]dto.Service)
		for _, service := range services {
			if service.Spec.Mode.Replicated.Replicas == 0 {
				continue
			}

			portStr, portKeyFound := service.Spec.TaskTemplate.ContainerSpec.Labels[LABEL_PORT];
			name, nameKeyFound := service.Spec.TaskTemplate.ContainerSpec.Labels[LABEL_NAME];

			if !portKeyFound || !nameKeyFound {
				continue
			}

			port, err := strconv.ParseInt(portStr, 10, 16)
			if err != nil {
				continue
			}

			_, serviceFound := servicesDto[name]
			if serviceFound {
				continue
			}

			alias := findAliasThatCanBeCalled(&detailedMe, &service)
			if alias == "" {
				continue
			}

			servicesDto[name] = dto.Service{
				Name: name,
				Port: uint16(port),
				Alias: alias,
			}
		}

		routes := dto.ServiceRoutes{
			ProviderName: PROVIDER_NAME,
			Services: servicesDto,
		}
		configurationChan <- routes

		sleepFunc()
	}
}

func findAliasThatCanBeCalled(containerCallee * docker.Container, serviceToBeCalled * docker.Service) string {
	for calleeNetworkName, _ := range containerCallee.NetworkSettings.Networks {
		for _,toBeCalledNetworkName := range serviceToBeCalled.Spec.TaskTemplate.Networks {
			if calleeNetworkName == toBeCalledNetworkName.Target && len(toBeCalledNetworkName.Aliases) > 0 {
				// Just pick one, no need to know exactly
				return toBeCalledNetworkName.Aliases[0]
			}
		}
	}

	return ""
}