package module

import (
	"collector/config.go"
	"collector/utils"
)

type ProcessConfig struct {
	ServiceInfo           config.ServiceConfig
	InstallationCommand   string
	InstallationArgs      []string
	UninstallationCommand string
	UninstallationArgs    []string
	ConnectionHost        string
	ConnectionKey         string
	Logger                *utils.BeautyLogger
}

type CollectorProcess interface {
	Run() error
	Install() error
	Uninstall() error
}

func GetCollectorProcess(collector config.Collector, host, key string, logger *utils.BeautyLogger) CollectorProcess {
	switch collector {
	case config.AS400:
		return getAS400Collector(host, key, logger)
	default:
		return nil
	}
}
