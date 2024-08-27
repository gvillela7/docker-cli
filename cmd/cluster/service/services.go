package service

import (
	"github.com/gvillela7/snake-docker-cli/cmd/config"
	"io"
	"log"
	"net/http"
)

func DockerNodes() []byte {
	dockerHost := config.ViperEnvVariable("DOCKER_HOST")
	response, err := http.Get(dockerHost + "/nodes")
	if err != nil {
		log.Fatalf("Get docker nodes failed . %s", err)
	}

	defer response.Body.Close()
	contents, _ := io.ReadAll(response.Body)

	return contents
}

func DockerServices() []byte {
	dockerHost := config.ViperEnvVariable("DOCKER_HOST")
	response, err := http.Get(dockerHost + "/services")
	if err != nil {
		log.Fatalf("Get docker services failed . %s", err)
	}

	defer response.Body.Close()
	contents, _ := io.ReadAll(response.Body)

	return contents
}
