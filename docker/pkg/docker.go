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

type Container struct {
	Id string
	State string
	Labels map[string]string
	NetworkSettings NetworkSettings
}

type Network struct {
	Aliases []string
}

type NetworkSettings struct {
	Networks map[string]Network
}

type DetailedContainer struct {
	Id string
	NetworkSettings NetworkSettings
}

type Docker interface {
	GetContainers() ([]Container, error)
	GetContainer(containerId string) (DetailedContainer, error)
	LinuxOnly_Me() (DetailedContainer, error)
}
