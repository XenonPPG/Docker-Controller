package controllers

import (
	compose "DockerController/gen/compose_v1"
	resourceusage "DockerController/gen/resources_messages_v1"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

const ProjectLabel = "com.docker.compose.project"

func (s *Server) CreateProject(ctx context.Context, req *compose.CreateProjectRequest) (*compose.Project, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) DeleteProject(ctx context.Context, req *compose.ProjectRequest) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) StopProject(ctx context.Context, req *compose.ProjectRequest) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) StartProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) RestartProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetProjectResourceUsage(ctx context.Context, req *compose.ProjectRequest) (*compose.GetProjectResourceUsageResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetProjectTotalResourceUsage(ctx context.Context, req *compose.ProjectRequest) (*resourceusage.ResourceUsage, error) {
	// TODO: implement
	return nil, nil
}
