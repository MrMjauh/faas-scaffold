package dto

type Service struct {
	Name string
}

type ServiceRoutes struct {
	ProviderName  string
	Services []Service
}