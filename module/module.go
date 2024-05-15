package module

import (
	"collector/config.go"
)

type ProcessConfig struct {
	ServiceInfo           config.ServiceConfig
	InstallationCommand   string
	InstallationArgs      []string
	UninstallationCommand string
	UninstallationArgs    []string
	ConnectionHost        string
	ConnectionKey         string
}

type CollectorProcess interface {
	Run() error
	Install() error
	Uninstall() error
}

func GetCollectorProcess(collector config.Collector, host, key string) CollectorProcess {
	switch collector {
	case config.AS400:
		return getAS400Collector(host, key)
	default:
		return nil
	}
}
