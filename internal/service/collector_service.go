package service

import (
	"collector/internal/config"
	"collector/internal/domain"
)

type CollectorService struct {
	Collector        domain.Collector
	Manager          domain.OSManager
	Host             string
	ConnectionKey    string
	CollectorConfigs *[]domain.CollectorConfig
	Templates        map[string]string
}

func NewCollectorService(manager domain.OSManager, collector domain.Collector, host string, connectionKey string) *CollectorService {
	return &CollectorService{
		Collector:        collector,
		Manager:          manager,
		Host:             host,
		ConnectionKey:    connectionKey,
		CollectorConfigs: config.LoadCollectorConfig("path/to/collectors.json"),
		Templates:        config.LoadServiceTemplates("path/to/service_templates.json"),
	}
}
func (c *CollectorService) SetupCollector() error {
	var selectedCollector domain.CollectorConfig

	getConfigHandler := &GetCollectorConfigHandler{}
	installDepsHandler := &InstallDependenciesHandler{}
	downloadCollectorHandler := &DownloadCollectorHandler{}
	createServiceHandler := &CreateServiceHandler{}

	getConfigHandler.SetNext(installDepsHandler)
	installDepsHandler.SetNext(downloadCollectorHandler)
	downloadCollectorHandler.SetNext(createServiceHandler)

	// Start execution
	if err := getConfigHandler.Execute(c, &selectedCollector); err != nil {
		return err
	}

	return nil
}
