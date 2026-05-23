package controllers

import (
	"DockerController/app/internal/initializers"
	"DockerController/app/internal/utils"
	compose "DockerController/gen/compose_v1"
	container "DockerController/gen/container_v1"
	dockercontroller "DockerController/gen/docker_controller_v1"
	resourceusage "DockerController/gen/resources_messages_v1"
	"context"
	"runtime"
	"strconv"

	"github.com/moby/moby/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func (s *Server) Ping(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *Server) GetResourceUsage(ctx context.Context, req *emptypb.Empty) (*dockercontroller.GetResourceUsageResponse, error) {
	containers, err := s.ListContainers(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to list containers: "+err.Error())
	}

	resourceMap := make(map[string]*dockercontroller.ResourceUsageMapValue)
	for _, c := range containers.Containers {
		res, err := s.GetContainerResourceUsage(ctx, &container.ContainerRequest{Id: c.GetId()})
		if err != nil {
			res = nil
		}
		resourceMap[c.GetId()] = &dockercontroller.ResourceUsageMapValue{
			Image:         c.GetImage(),
			ResourceUsage: res,
		}
	}

	return &dockercontroller.GetResourceUsageResponse{
		ContainerUsage: resourceMap,
	}, nil
}

func (s *Server) GetTotalResourceUsage(ctx context.Context, req *emptypb.Empty) (*resourceusage.ResourceUsage, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get cpu stats: "+err.Error())
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get memory stats: "+err.Error())
	}

	netStats, err := net.IOCounters(false) // false = суммарно по всем интерфейсам
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get network stats: "+err.Error())
	}

	diskStats, err := disk.IOCounters()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get disk stats: "+err.Error())
	}

	// суммируем все диски
	var totalRead, totalWrite uint64
	for _, d := range diskStats {
		totalRead += d.ReadBytes
		totalWrite += d.WriteBytes
	}

	networks := make([]*resourceusage.NetworkStats, 0, len(netStats))
	for _, n := range netStats {
		networks = append(networks, &resourceusage.NetworkStats{
			Name:    n.Name,
			RxBytes: n.BytesRecv,
			TxBytes: n.BytesSent,
		})
	}

	return &resourceusage.ResourceUsage{
		Cpu: &resourceusage.CpuStats{
			Percent:    cpuPercent[0],
			OnlineCpus: int32(runtime.NumCPU()),
		},
		Memory: &resourceusage.MemoryStats{
			Usage:   memStat.Used,
			Limit:   memStat.Total,
			Percent: memStat.UsedPercent,
		},
		Networks: networks,
		BlockIo: &resourceusage.BlockIOStats{
			ReadBytes:  totalRead,
			WriteBytes: totalWrite,
		},
	}, nil
}

func (s *Server) ListContainers(ctx context.Context, req *emptypb.Empty) (*dockercontroller.ListContainersResponse, error) {
	resp, err := initializers.DockerClient.ContainerList(ctx, client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	var containers = make([]*container.Container, len(resp.Items))
	for i, s := range resp.Items {
		containers[i] = &container.Container{
			Id:        s.ID,
			Name:      s.Names[0][:1],
			Image:     s.Image,
			Status:    s.Status,
			CreatedAt: strconv.FormatInt(s.Created, 10),
			Labels:    s.Labels,
		}
	}

	return &dockercontroller.ListContainersResponse{
		Containers: containers,
	}, nil
}

func (s *Server) ListProjects(ctx context.Context, req *emptypb.Empty) (*dockercontroller.ListProjectsResponse, error) {
	resp, err := initializers.DockerClient.ContainerList(ctx, client.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	labelsMap := make(map[string][]*container.Container)
	for _, s := range resp.Items {
		containers, ok := labelsMap[s.Labels[ProjectLabel]]
		if !ok {
			containers = make([]*container.Container, 0)
		}
		containers = append(containers, &container.Container{
			Id:        s.ID,
			Name:      s.Names[0][:1],
			Image:     s.Image,
			Status:    s.Status,
			CreatedAt: strconv.FormatInt(s.Created, 10),
		})
		labelsMap[s.Labels[ProjectLabel]] = containers
	}

	projects := make([]*compose.Project, 0, len(labelsMap))
	for k, v := range labelsMap {
		projects = append(projects, &compose.Project{
			Name:       k,
			Status:     utils.ProjectStatus(v),
			Containers: v,
		})
	}

	return &dockercontroller.ListProjectsResponse{
		Projects: projects,
	}, nil
}
