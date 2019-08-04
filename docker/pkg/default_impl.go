package docker

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
)

type DefaultImpl struct {

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

func (defaultImpl DefaultImpl) ListAllContainers() ([]Container, error) {
	client := createClient()
	resp, err := client.Get("http://unix" + "/containers/json")

	if err != nil {
		return []Container{}, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	var containers []Container
	json.Unmarshal([]byte(body), &containers)

	return containers, nil
}