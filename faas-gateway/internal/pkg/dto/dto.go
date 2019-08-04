package dto

type Service struct {
	Name string
	Port uint16
}

type ServiceRoutes struct {
	ProviderName string
	Services map[string]Service
}