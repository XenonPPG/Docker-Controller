package utils

import (
	container "DockerController/gen/container_v1"
	"strings"
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
