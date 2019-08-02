package docker

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"encoding/json"
)

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
}

func createClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}
}

func ListAllContainers() []Container {
	client := createClient()
	resp, err := client.Get("http://unix" + "/containers/json")

	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	var containers []Container
	json.Unmarshal([]byte(body), &containers)

	return containers
}