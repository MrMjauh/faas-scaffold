package main

import (
	docker "faas-scaffold/docker/pkg"
	"fmt"
)

func main() {
	docker.ListAllContainers()
	fmt.Println("Hello Gateway")
}
