package controllers

import (
	"DockerController/app/internal/initializers"
	container "DockerController/gen/container_v1"
	resourceusage "DockerController/gen/resources_messages_v1"
	"context"
	"encoding/json"
	"io"
	"log"

	types "github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) CreateContainer(ctx context.Context, req *container.CreateContainerRequest) (*container.Container, error) {
	image := req.GetImage()

	// pull image if it's not present
	_, err := initializers.DockerClient.ImageInspect(ctx, image)
	if err != nil {
		reader, err := initializers.DockerClient.ImagePull(ctx, image, client.ImagePullOptions{})
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to pull image: "+err.Error())
		}
		defer func() {
			_ = reader.Close()
		}()

		if _, err := io.Copy(io.Discard, reader); err != nil {
			return nil, status.Error(codes.Internal, "failed reading image pull response: "+err.Error())
		}
	}

	result, err := initializers.DockerClient.ContainerCreate(
		ctx,
		client.ContainerCreateOptions{
			Name:  req.GetName(),
			Image: image,
		},
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create container: "+err.Error())
	}

	return s.GetContainer(ctx, &container.ContainerRequest{Id: result.ID})
}

func (s *Server) GetContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	if initializers.DockerClient == nil {
		return nil, status.Error(codes.Internal, "docker client is not initialized")
	}

	info, err := initializers.DockerClient.ContainerInspect(ctx, req.GetId(), client.ContainerInspectOptions{})
	if err != nil {
		return nil, err
	}

	cInfo := info.Container

	containerStatus := "unknown"
	if cInfo.State != nil {
		containerStatus = string(cInfo.State.Status)
	}

	return &container.Container{
		Id:        cInfo.ID,
		Name:      cInfo.Name,
		Image:     cInfo.Image,
		Status:    containerStatus,
		CreatedAt: cInfo.Created,
	}, nil
}

func (s *Server) DeleteContainer(ctx context.Context, req *container.ContainerRequest) (*emptypb.Empty, error) {
	// container remove results are not implemented in the current client version
	_, err := initializers.DockerClient.ContainerRemove(ctx, req.GetId(), client.ContainerRemoveOptions{
		Force: true,
	})

	_, err = s.CleanUp(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("failed to clean up images: ", err)
	}

	return &emptypb.Empty{}, err
}

func (s *Server) StartContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	cont, err := s.GetContainer(ctx, req)
	if err == nil && cont != nil {
		switch cont.Status {
		case "running":
			return cont, nil
		case "paused":
			// container unpause results are not implemented in the client yet
			_, err = initializers.DockerClient.ContainerUnpause(ctx, req.GetId(), client.ContainerUnpauseOptions{})

			// if container was successfully unpaused, break the function
			if err == nil {
				return s.GetContainer(ctx, req)
			}
		}
	}

	// container start results are not implemented in the client yet
	_, err = initializers.DockerClient.ContainerStart(ctx, req.GetId(), client.ContainerStartOptions{})

	return s.GetContainer(ctx, req)
}

func (s *Server) RestartContainer(ctx context.Context, req *container.ContainerRequest) (*container.Container, error) {
	// container restart results are not implemented in the current client version
	_, err := initializers.DockerClient.ContainerRestart(ctx, req.GetId(), client.ContainerRestartOptions{})
	return nil, err
}

func (s *Server) PauseContainer(ctx context.Context, req *container.ContainerRequest) (*emptypb.Empty, error) {
	// container pause results are not implemented in the current client version
	_, err := initializers.DockerClient.ContainerPause(ctx, req.GetId(), client.ContainerPauseOptions{})
	return &emptypb.Empty{}, err
}

func (s *Server) StopContainer(ctx context.Context, req *container.ContainerRequest) (*emptypb.Empty, error) {
	_, err := initializers.DockerClient.ContainerStop(ctx, req.GetId(), client.ContainerStopOptions{})
	return &emptypb.Empty{}, err
}

func (s *Server) GetContainerResourceUsage(ctx context.Context, req *container.ContainerRequest) (*resourceusage.ResourceUsage, error) {
	stats, err := initializers.DockerClient.ContainerStats(ctx, req.GetId(), client.ContainerStatsOptions{
		Stream:                false,
		IncludePreviousSample: true,
	})
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stats.Body.Close()
	}()

	var resp types.StatsResponse
	if err = json.NewDecoder(stats.Body).Decode(&resp); err != nil {
		return nil, err
	}

	// cpu stats
	cpuDelta := float64(resp.CPUStats.CPUUsage.TotalUsage - resp.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(resp.CPUStats.SystemUsage - resp.PreCPUStats.SystemUsage)
	numCPUs := resp.CPUStats.OnlineCPUs
	cpuPercent := (cpuDelta / systemDelta) * float64(numCPUs) * 100.0
	CPUStats := &resourceusage.CpuStats{
		Percent:    cpuPercent,
		OnlineCpus: int32(numCPUs),
	}

	// memory stats
	memUsage := resp.MemoryStats.Usage - resp.MemoryStats.Stats["cache"]
	memLimit := resp.MemoryStats.Limit
	memPercent := float64(memUsage) / float64(memLimit) * 100.0
	memStats := &resourceusage.MemoryStats{
		Usage:   memUsage,
		Limit:   memLimit,
		Percent: memPercent,
	}

	// network stats
	networks := make([]*resourceusage.NetworkStats, 0, len(resp.Networks))
	for name, network := range resp.Networks {
		networks = append(networks, &resourceusage.NetworkStats{
			Name:    name,
			RxBytes: network.RxBytes,
			TxBytes: network.TxBytes,
		})
	}

	// block io stats
	blockIoStats := &resourceusage.BlockIOStats{
		ReadBytes:  0,
		WriteBytes: 0,
	}
	for _, ioStat := range resp.BlkioStats.IoServiceBytesRecursive {
		switch ioStat.Op {
		case "Read":
			blockIoStats.ReadBytes += ioStat.Value
		case "Write":
			blockIoStats.WriteBytes += ioStat.Value
		}
	}

	result := &resourceusage.ResourceUsage{
		Cpu:      CPUStats,
		Memory:   memStats,
		Networks: networks,
		BlockIo:  blockIoStats,
	}

	return result, nil
}
