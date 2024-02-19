package infra

import (
	"collector/internal/domain"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type LinuxManager struct{}

func (lm *LinuxManager) CheckAndInstallDependencies(dependencies []domain.Dependency) error {
	// Example: Check if Java is installed; if not, install it using apt-get for Ubuntu/Debian.
	for _, dep := range dependencies {
		fmt.Printf("Checking if %s is installed...\n", dep.Name)
		if _, err := exec.LookPath(dep.CheckCommand); err != nil {
			fmt.Printf("%s is not installed, installing...\n", dep.Name)
			installCmd := exec.Command("sh", "-c", dep.InstallCommand["linux"])
			if err := installCmd.Run(); err != nil {
				return fmt.Errorf("failed to install %s: %v", dep.Name, err)
			}
		}
	}
	return nil
}

func (lm *LinuxManager) DownloadCollector(collector domain.CollectorConfig) (string, error) {
	// Download the collector using HTTP GET request
	resp, err := http.Get(collector.DownloadUrl)
	if err != nil {
		return "", fmt.Errorf("failed to download collector: %v", err)
	}
	defer resp.Body.Close()

	// Create a file in /tmp (or another directory, adjust as needed)
	fileName := filepath.Base(collector.DownloadUrl)
	filePath := filepath.Join("/tmp", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file for collector: %v", err)
	}
	defer file.Close()

	// Write the response body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to write collector to file: %v", err)
	}

	return filePath, nil
}

func (lm *LinuxManager) CreateService(config domain.CollectorConfig, servicePath string, template string) error {
	// Prepare the service file content by replacing placeholders in the template
	// This example assumes template is a path to the template file, and placeholders are replaced elsewhere
	serviceContent, err := os.ReadFile(template)
	if err != nil {
		return fmt.Errorf("failed to read service template: %v", err)
	}

	// This is a simplification. In practice, you'd replace placeholders in serviceContent with actual values

	// Write the service file to /etc/systemd/system
	serviceName := fmt.Sprintf("%s.service", config.Name)
	serviceFilePath := filepath.Join("/etc/systemd/system", serviceName)
	if err := os.WriteFile(serviceFilePath, serviceContent, 0644); err != nil {
		return fmt.Errorf("failed to write service file: %v", err)
	}

	// Reload systemd manager configuration
	if err := exec.Command("systemctl", "daemon-reload").Run(); err != nil {
		return fmt.Errorf("failed to reload systemd daemon: %v", err)
	}

	// Enable and start the service
	if err := exec.Command("systemctl", "enable", serviceName).Run(); err != nil {
		return fmt.Errorf("failed to enable service: %v", err)
	}
	if err := exec.Command("systemctl", "start", serviceName).Run(); err != nil {
		return fmt.Errorf("failed to start service: %v", err)
	}

	return nil
}

func (lm *LinuxManager) DisableService(serviceName string) error {
	// Stop the service
	if err := exec.Command("systemctl", "stop", serviceName+".service").Run(); err != nil {
		return fmt.Errorf("failed to stop service: %v", err)
	}
	// Disable the service
	if err := exec.Command("systemctl", "disable", serviceName+".service").Run(); err != nil {
		return fmt.Errorf("failed to disable service: %v", err)
	}
	return nil
}

// Ensure LinuxManager implements domain.OSManager
var _ domain.OSManager = &LinuxManager{}
