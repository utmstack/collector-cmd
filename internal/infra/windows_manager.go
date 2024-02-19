package infra

import (
	"collector/internal/domain"
	"fmt"
	"os/exec"
)

type WindowsManager struct{}

func (wm *WindowsManager) CheckAndInstallDependencies(dependencies []domain.Dependency) error {
	for _, dep := range dependencies {
		// Check if the dependency is installed
		checkCmd := exec.Command("powershell", "-Command", dep.CheckCommand)
		if err := checkCmd.Run(); err != nil {
			fmt.Printf("%s is not installed. Installing...\n", dep.Name)
			// Install the dependency using PowerShell
			installCmdStr := dep.InstallCommand["windows"]
			installCmd := exec.Command("powershell", "-Command", installCmdStr)
			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("failed to install %s: %v", dep.Name, err)
			}
		}
	}
	return nil
}

func (wm *WindowsManager) DownloadCollector(collector domain.CollectorConfig) (string, error) {
	// Implement the logic to download the collector based on the provided URL.
	// For simplicity, this example just returns a placeholder path.
	return "C:\\path\\to\\collector.exe", nil
}

func (wm *WindowsManager) CreateService(config domain.CollectorConfig, servicePath string, template string) error {
	// Implement the logic to create a Windows service based on the downloaded collector.
	// This is a placeholder. You'll need to use Windows APIs or external tools like NSSM, sc.exe, etc.
	fmt.Println("Creating Windows service for collector:", config.Name)
	return nil
}

func (wm *WindowsManager) DisableService(serviceName string) error {
	// Stop the service
	if err := exec.Command("sc", "stop", serviceName).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %v", err)
	}
	// Disable the service
	if err := exec.Command("sc", "config", serviceName, "start=", "disabled").Run(); err != nil {
		return fmt.Errorf("failed to disable service: %v", err)
	}
	return nil
}

var _ domain.OSManager = &WindowsManager{}
