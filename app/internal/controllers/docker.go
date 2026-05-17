package controllers

import (
	dockercontroller "DockerController/gen/docker_controller_v1"
	"DockerController/gen/resources_messages_v1"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetInfo(ctx context.Context, req *emptypb.Empty) (*dockercontroller.DockerInfo, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetVersion(ctx context.Context, req *emptypb.Empty) (*dockercontroller.DockerVersion, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) GetResourceUsage(ctx context.Context, req *emptypb.Empty) (*resources_messages_v1.ResourceUsage, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) StreamEvents(req *dockercontroller.StreamEventsRequest, stream dockercontroller.DockerService_StreamEventsServer) error {
	// TODO: implement
	return nil
}

func (s *Server) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	// TODO: implement
	return nil, nil
}
