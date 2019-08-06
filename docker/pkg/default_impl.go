package docker

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

type DefaultImpl struct {

}

const (
	VERSION = "v1.24"
)

func createClient() http.Client {
	return http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}
}

func (defaultImpl DefaultImpl) GetServices() ([]Service, error) {
	client := createClient()
	resp, err := client.Get("http://unix" + "/" + VERSION + "/services")

	if err != nil {
		return []Service{}, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	var services []Service
	json.Unmarshal([]byte(body), &services)

	return services, nil
}

func (defaultImpl DefaultImpl) GetContainer(containerId string) (Container, error) {
	client := createClient()
	resp, err := client.Get("http://unix/" + VERSION + "/containers/" + containerId + "/json")

	if err != nil {
		return Container{}, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	var container Container
	json.Unmarshal([]byte(body), &container)

	return container, nil
}

func (defaultImpl DefaultImpl) LinuxOnly_Me() (Container, error) {
	out, err := exec.Command("cat", "/etc/hostname").Output()
	if err != nil {
		log.Println(err)
		return Container{}, err
	}
	containerId := strings.TrimSuffix(strings.TrimSpace(string(out)), "\n")

	container, err := defaultImpl.GetContainer(containerId)
	return container, nil
}