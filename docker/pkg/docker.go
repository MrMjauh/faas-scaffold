package docker

const (
	STATUS_STATE_CREATED = "created"
	STATUS_STATE_RESTARTING = "restarting"
	STATUS_STATE_RUNNING = "running"
	STATUS_STATE_REMOVING = "removing"
	STATUS_STATE_PAUSED = "paused"
	STATUS_STATE_EXITED = "exited"
	STATUS_STATE_DEAD = "dead"
)

// Example of a service json from docker api 1.24
// {"ID":"dcbsjw0mpqzuuujg8m3fjiylu","Version":{"Index":1014},"CreatedAt":"2019-08-04T20:59:50.388523183Z","UpdatedAt":"2019-08-05T15:40:08.539319236Z","Spec":{"Name":"faas_faas-gateway","Labels":{"com.docker.stack.image":"faas-scaffold_faas-gateway:latest","com.docker.stack.namespace":"faas"},"TaskTemplate":{"ContainerSpec":{"Image":"faas-scaffold_faas-gateway:latest","Labels":{"com.docker.stack.namespace":"faas"},"Privileges":{"CredentialSpec":null,"SELinuxContext":null},"Mounts":[{"Type":"bind","Source":"/var/run/docker.sock","Target":"/var/run/docker.sock"}],"Isolation":"default"},"Resources":{},"Placement":{},"Networks":[{"Target":"3o7191d94rmd5amoc3pciec38","Aliases":["faas-gateway"]}],"ForceUpdate":0,"Runtime":"container"},"Mode":{"Replicated":{"Replicas":1}},"EndpointSpec":{"Mode":"vip","Ports":[{"Protocol":"tcp","TargetPort":8081,"PublishedPort":8081,"PublishMode":"ingress"}]}},"Endpoint":{"Spec":{"Mode":"vip","Ports":[{"Protocol":"tcp","TargetPort":8081,"PublishedPort":8081,"PublishMode":"ingress"}]},"Ports":[{"Protocol":"tcp","TargetPort":8081,"PublishedPort":8081,"PublishMode":"ingress"}],"VirtualIPs":[{"NetworkID":"pn5zzhrge631xkv1wr42ewb89","Addr":"10.255.0.13/16"},{"NetworkID":"3o7191d94rmd5amoc3pciec38","Addr":"10.0.9.6/24"}]}}

type ServiceNetwork struct {
	Target string
	Aliases []string
}

type ServiceReplicated struct {
	Replicas int
}
type ServiceMode struct {
	Replicated ServiceReplicated
}

type ServiceContainerSpec struct {
	Labels map[string]string
}

type ServiceTaskTemplate struct {
	Networks []ServiceNetwork
	ContainerSpec ServiceContainerSpec
}

type ServiceSpec struct {
	Name string
	TaskTemplate ServiceTaskTemplate
	Mode ServiceMode
}

type Service struct {
	ID string
	Spec ServiceSpec
}

type Container struct {
	Id string
	NetworkSettings ContainerNetworkSettings
}

type ContainerNetwork struct {
	Aliases []string
}

type ContainerNetworkSettings struct {
	Networks map[string]ContainerNetwork
}

type Docker interface {
	GetServices() ([]Service, error)
	GetContainer(containerId string) (Container, error)
	LinuxOnly_Me() (Container, error)
}
