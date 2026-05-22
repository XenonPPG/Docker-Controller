package initializers

import (
	"log"

	"github.com/moby/moby/client"
)

var DockerClient *client.Client

func InitDockerClient() *client.Client {
	var err error

	DockerClient, err = client.New(client.FromEnv)
	if err != nil {
		panic(err)
	}

	log.Println("Docker client initialized")

	return DockerClient
}
