package service

import (
	"collector/internal/domain"
	"collector/internal/infra"
	"runtime"
)

func NewOSManager() domain.OSManager {
	switch runtime.GOOS {
	case "windows":
		return &infra.WindowsManager{}
	case "linux":
		return &infra.LinuxManager{}
	default:
		return &infra.LinuxManager{}
	}
}
