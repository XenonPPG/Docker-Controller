package controllers

import (
	container "DockerController/gen/container_v1"
	resourceusage "DockerController/gen/resources_messages_v1"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateContainer(ctx context.Context, req *container.CreateContainerRequest) (*container.Container, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) DeleteContainer(ctx context.Context, req *container.ContainerRequest) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) StartContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) RestartContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) StopContainer(ctx context.Context, req *container.ContainerRequest) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetContainerResourceUsage(ctx context.Context, req *container.ContainerRequest) (*resourceusage.ResourceUsage, error) {
	// TODO: implement
	return nil, nil
}
