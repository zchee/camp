package utils

import (
	"net/url"
	"os"

	log "github.com/Sirupsen/logrus"
)

func GetDockerHost() string {
	dockerTcpHost := os.Getenv("DOCKER_HOST")
	u, err := url.Parse(dockerTcpHost)
	if err != nil {
		panic(err)
	}
	if dockerTcpHost == "" {
		log.Fatalf("$DOCKER_HOST not found.")
	}
	dockerHost := u.Host

	return dockerHost
}
