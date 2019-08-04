package dto

type Service struct {
	Name string
	Port uint16
	Alias string
}

type ServiceRoutes struct {
	ProviderName string
	Services map[string]Service
}