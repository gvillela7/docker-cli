package service

import (
	"github.com/gvillela7/snake-docker-cli/cmd/config"
	"io"
	"log"
	"net/http"
)

func ContainersList() []byte {
	dockerHost := config.ViperEnvVariable("DOCKER_HOST")
	response, err := http.Get(dockerHost + "/containers/json")
	if err != nil {
		log.Fatalf("Get docker containers failed . %s", err)
	}

	defer response.Body.Close()
	contents, _ := io.ReadAll(response.Body)

	return contents
}

func GetContainerByID(id string) []byte {
	dockerHost := config.ViperEnvVariable("DOCKER_HOST")
	response, err := http.Get(dockerHost + "/containers/" + id + "/json")
	if err != nil {
		log.Fatalf("Inspect container failed . %s", err)
	}

	defer response.Body.Close()
	contents, _ := io.ReadAll(response.Body)

	return contents
}
