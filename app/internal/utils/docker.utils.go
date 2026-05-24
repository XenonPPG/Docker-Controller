package utils

import (
	container "DockerController/gen/container_v1"
	"errors"
	"strconv"
	"strings"
	"sync"

	clientcontainer "github.com/moby/moby/api/types/container"
)

const (
	ProjectStatusUnknown = "unknown"
	ProjectStatusRunning = "running"
	ProjectStatusPartial = "partial"
	ProjectStatusStopped = "stopped"
)

func ProjectStatus(containers []*container.Container) string {
	status := ProjectStatusUnknown

	anyRunning := false
	allRunning := true
	for _, c := range containers {
		running := strings.HasPrefix(c.Status, "Up")
		if running {
			anyRunning = true
		} else {
			allRunning = false
			if anyRunning {
				break
			}
		}
	}

	if allRunning {
		status = ProjectStatusRunning
	} else if anyRunning {
		status = ProjectStatusPartial
	} else {
		status = ProjectStatusStopped
	}

	return status
}

func MapSummaryToContainer(summary clientcontainer.Summary) *container.Container {
	return &container.Container{
		Id:        summary.ID,
		Name:      summary.Names[0][:1],
		Image:     summary.Image,
		Status:    summary.Status,
		CreatedAt: strconv.FormatInt(summary.Created, 10),
	}
}

func ProcessContainers(containers []*container.Container, f func(c *container.Container) error) error {
	wg := sync.WaitGroup{}
	wg.Add(len(containers))

	errCh := make(chan error, len(containers))

	for _, c := range containers {
		go func() {
			defer wg.Done()

			if err := f(c); err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)

	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}
