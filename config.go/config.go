package config

import (
	"collector/utils"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type ServiceTypeConfig struct {
	CollectorType Collector `yaml:"collector_type"`
}

const (
	As400Jar            = "as400-extractor-2.2.0-jar-with-dependencies.jar"
	JavaVersion         = "jdk-11.0.13.8"
	CollectorConfigFile = "collector-config.yaml"
	SERV_LOG            = "utmstack_collector.log"
)

type ServiceConfig struct {
	Name        string
	DisplayName string
	Description string
	CMDRun      string
	CMDArgs     []string
	CMDPath     string
}

type Collector string

var (
	AS400 Collector = "as400"
)

func IsValidCollector(c string) bool {
	switch c {
	case string(AS400):
		return true
	default:
		return false
	}
}

func GetJavaPath() string {
	currentPath, _ := utils.GetMyPath()
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(currentPath, "dependencies", JavaVersion, "bin", "java.exe")
	default:
		return ""
	}
}

func GetAs400Command(action, host, connectionKey string) (string, []string) {
	currentPath, _ := utils.GetMyPath()
	cmd := ""
	args := []string{}

	action = strings.ToUpper(action)
	switch runtime.GOOS {
	case "windows":
		cmd = GetJavaPath()
	case "linux":
		cmd = "java"
	default:
		return "unknown operating system", nil
	}

	args = append(args, "-jar", filepath.Join(currentPath, As400Jar), fmt.Sprintf("-option=%s", action))
	if action == "INSTALL" {
		args = append(args, fmt.Sprintf("-collector-manager-host=%s", host), "-collector-manager-port=9000", "-logs-port=50051", fmt.Sprintf("-connection-key=%s", connectionKey))
	}
	return cmd, args
}

func ValidateParams(params []string) bool {
	action := params[1]
	collector := params[2]
	if !IsValidCollector(collector) {
		return false
	}

	if action == "install" {
		if len(params) < 5 {
			return false
		}
	} else if action == "uninstall" {
		if len(params) < 3 {
			return false
		}
	} else {
		return false
	}

	return true
}

func GetServiceConfig(c Collector) ServiceConfig {
	path, _ := utils.GetMyPath()
	switch c {
	case AS400:
		cmd, args := GetAs400Command("run", "", "")
		return ServiceConfig{
			Name:        "UTMStackAS400Collector",
			DisplayName: "UTMStack AS400 Collector",
			Description: "UTMStack AS400 Collector",
			CMDRun:      cmd,
			CMDArgs:     args,
			CMDPath:     path,
		}
	default:
		return ServiceConfig{}
	}
}

func SaveConfig(config *ServiceTypeConfig) error {
	path, err := utils.GetMyPath()
	if err != nil {
		return fmt.Errorf("failed to get current path: %v", err)
	}

	if err := utils.WriteYAML(filepath.Join(path, CollectorConfigFile), config); err != nil {
		return err
	}

	return nil
}

func ReadConfig() (*ServiceTypeConfig, error) {
	path, err := utils.GetMyPath()
	if err != nil {
		return nil, fmt.Errorf("failed to get current path: %v", err)
	}

	config := &ServiceTypeConfig{}
	if err = utils.ReadYAML(filepath.Join(path, CollectorConfigFile), config); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	return config, nil
}
