package controllers

import (
	"DockerController/app/internal/initializers"
	"DockerController/app/internal/utils"
	compose "DockerController/gen/compose_v1"
	container "DockerController/gen/container_v1"
	resourceusage "DockerController/gen/resources_messages_v1"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/moby/moby/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const ProjectLabel = "com.docker.compose.project"

func (s *Server) CreateProject(ctx context.Context, req *compose.CreateProjectRequest) (*compose.Project, error) {
	if err := os.MkdirAll("tmp", 0755); err != nil {
		return nil, err
	}

	dir, err := os.MkdirTemp("tmp", "compose-*")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := os.RemoveAll(dir); err != nil {
			log.Println("failed to remove temp dir:", err)
		}
	}()

	composePath := filepath.Join(dir, "docker-compose.yml")
	if err = os.WriteFile(composePath, []byte(req.GetComposeFile()), 0644); err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(ctx, "docker", "compose", "-f", composePath, "-p", req.GetName(), "up", "-d")
	if out, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("compose up failed: %w\noutput: %s", err, out)
	}

	return s.GetProject(ctx, &compose.ProjectRequest{Name: req.GetName()})
}

func (s *Server) GetProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	filters := make(client.Filters).Add("label", ProjectLabel+"="+req.GetName())

	filtered, err := initializers.DockerClient.ContainerList(ctx, client.ContainerListOptions{
		All:     true,
		Filters: filters,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to list containers: "+err.Error())
	}

	containers := make([]*container.Container, 0, len(filtered.Items))
	for _, c := range filtered.Items {
		containers = append(containers, utils.MapSummaryToContainer(c))
	}

	return &compose.Project{
		Name:       req.GetName(),
		Status:     utils.ProjectStatus(containers),
		Containers: containers,
	}, nil
}

func GenericProjectFunction(s *Server, ctx context.Context, req *compose.ProjectRequest, containerProcessor func(c *container.Container) error, description string) (*compose.Project, error) {
	project, err := s.GetProject(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get project: "+err.Error())
	}
	if project == nil || project.Containers == nil || len(project.Containers) == 0 {
		return project, nil
	}

	err = utils.ProcessContainers(project.Containers, containerProcessor)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to %s: %s", description, err.Error()))
	}
	return project, nil
}

func (s *Server) DeleteProject(ctx context.Context, req *compose.ProjectRequest) (*emptypb.Empty, error) {
	_, err := GenericProjectFunction(s, ctx, req, func(c *container.Container) error {
		_, err := s.DeleteContainer(ctx, &container.ContainerRequest{Id: c.Id})
		return err
	}, "delete project")

	_, err = s.CleanUp(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("failed to clean up images: ", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) StopProject(ctx context.Context, req *compose.ProjectRequest) (*emptypb.Empty, error) {
	_, err := GenericProjectFunction(s, ctx, req, func(c *container.Container) error {
		_, err := s.StopContainer(ctx, &container.ContainerRequest{Id: c.Id})
		return err
	}, "stop project")
	return &emptypb.Empty{}, err
}

func (s *Server) StartProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	project, err := GenericProjectFunction(s, ctx, req, func(c *container.Container) error {
		_, err := s.StartContainer(ctx, &container.ContainerRequest{Id: c.Id})
		return err
	}, "start project")
	return project, err
}

func (s *Server) RestartProject(ctx context.Context, req *compose.ProjectRequest) (*compose.Project, error) {
	project, err := GenericProjectFunction(s, ctx, req, func(c *container.Container) error {
		_, err := s.RestartContainer(ctx, &container.ContainerRequest{Id: c.Id})
		return err
	}, "restart project")
	return project, err
}

func (s *Server) GetProjectResourceUsage(ctx context.Context, req *compose.ProjectRequest) (*compose.GetProjectResourceUsageResponse, error) {
	project, err := s.GetProject(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get project: "+err.Error())
	}

	containers := project.GetContainers()
	if containers == nil || len(containers) == 0 {
		return nil, status.Error(codes.Internal, "no containers in project")
	}

	waitG := sync.WaitGroup{}
	waitG.Add(len(containers))

	resourceMap := make(map[string]*resourceusage.ResourceUsageMapValue)
	mu := sync.Mutex{}
	for _, c := range containers {
		go func() {
			res, err := s.GetContainerResourceUsage(ctx, &container.ContainerRequest{Id: c.GetId()})
			if err != nil {
				res = nil
			}
			mu.Lock()
			resourceMap[c.GetId()] = &resourceusage.ResourceUsageMapValue{
				Image:         c.GetImage(),
				ResourceUsage: res,
			}
			mu.Unlock()
			waitG.Done()
		}()
	}

	waitG.Wait()

	return &compose.GetProjectResourceUsageResponse{
		Stats: resourceMap,
	}, nil
}

func (s *Server) GetProjectTotalResourceUsage(ctx context.Context, req *compose.ProjectRequest) (*resourceusage.ResourceUsage, error) {
	project, err := s.GetProject(ctx, req)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get project: "+err.Error())
	}

	containers := project.GetContainers()
	if containers == nil || len(containers) == 0 {
		return nil, status.Error(codes.Internal, "no containers in project")
	}

	waitG := sync.WaitGroup{}
	waitG.Add(len(containers))

	totalUsage := &resourceusage.ResourceUsage{
		Cpu: &resourceusage.CpuStats{
			Percent:    0,
			OnlineCpus: 0,
		},
		Memory: &resourceusage.MemoryStats{
			Usage:   0,
			Limit:   0,
			Percent: 0,
		},
		Networks: []*resourceusage.NetworkStats{
			{
				Name:    "total",
				RxBytes: 0,
				TxBytes: 0,
			},
		},
		BlockIo: &resourceusage.BlockIOStats{
			ReadBytes:  0,
			WriteBytes: 0,
		},
	}

	filterNaN := func(n float64) float64 {
		if math.IsNaN(n) {
			return 0
		}
		return n
	}

	for _, c := range containers {
		go func() {
			defer waitG.Done()

			res, err := s.GetContainerResourceUsage(ctx, &container.ContainerRequest{Id: c.GetId()})
			if err != nil {
				res = nil
			}
			totalUsage.Cpu.Percent += filterNaN(res.GetCpu().GetPercent())
			totalUsage.Cpu.OnlineCpus = max(totalUsage.Cpu.OnlineCpus, res.GetCpu().OnlineCpus)
			totalUsage.Memory.Usage += res.GetMemory().GetUsage()
			totalUsage.Memory.Limit = max(totalUsage.Memory.Limit, res.GetMemory().GetLimit())
			totalUsage.Memory.Percent += filterNaN(res.GetMemory().GetPercent())
			for _, n := range res.GetNetworks() {
				totalUsage.Networks[0].RxBytes += n.GetRxBytes()
				totalUsage.Networks[0].TxBytes += n.GetTxBytes()
			}
			totalUsage.BlockIo.ReadBytes += res.GetBlockIo().GetReadBytes()
			totalUsage.BlockIo.WriteBytes += res.GetBlockIo().GetWriteBytes()
		}()
	}

	waitG.Wait()

	return totalUsage, nil
}
