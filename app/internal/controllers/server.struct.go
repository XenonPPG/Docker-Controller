package controllers

import (
	compose "DockerController/gen/compose_v1"
	container "DockerController/gen/container_v1"
	dockercontroller "DockerController/gen/docker_controller_v1"

	"google.golang.org/grpc"
)

type Server struct {
	dockercontroller.UnimplementedDockerServiceServer
	compose.UnimplementedComposeServiceServer
	container.UnimplementedContainerServiceServer
}

func (s *Server) RegisterFunc() func(server *grpc.Server) {
	return func(server *grpc.Server) {
		dockercontroller.RegisterDockerServiceServer(server, s)
		compose.RegisterComposeServiceServer(server, s)
		container.RegisterContainerServiceServer(server, s)
	}
}
