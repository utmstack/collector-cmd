package service

import (
	"collector/internal/domain"
)

type Handler interface {
	Execute(c *CollectorService, selectedCollector *domain.CollectorConfig) error
	SetNext(handler Handler)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

type GetCollectorConfigHandler struct {
	BaseHandler
}

func (h *GetCollectorConfigHandler) Execute(c *CollectorService, selectedCollector *domain.CollectorConfig) error {
	for _, col := range *c.CollectorConfigs {
		if col.Name == c.Collector {
			*selectedCollector = col
			break
		}
	}
	if h.next != nil {
		return h.next.Execute(c, selectedCollector)
	}
	return nil
}

// InstallDependenciesHandler checks and installs necessary dependencies for the collector.
type InstallDependenciesHandler struct {
	BaseHandler
}

func (h *InstallDependenciesHandler) Execute(c *CollectorService, selectedCollector *domain.CollectorConfig) error {
	if err := c.Manager.CheckAndInstallDependencies(selectedCollector.Dependencies); err != nil {
		return err
	}
	if h.next != nil {
		return h.next.Execute(c, selectedCollector)
	}
	return nil
}

// DownloadCollectorHandler downloads the collector based on the configuration.
type DownloadCollectorHandler struct {
	BaseHandler
}

func (h *DownloadCollectorHandler) Execute(c *CollectorService, selectedCollector *domain.CollectorConfig) error {
	downloadedPath, err := c.Manager.DownloadCollector(*selectedCollector)
	if err != nil {
		return err
	}
	selectedCollector.DownloadUrl = downloadedPath // Assuming CollectorConfig has a DownloadedPath field to store this.
	if h.next != nil {
		return h.next.Execute(c, selectedCollector)
	}
	return nil
}

// CreateServiceHandler sets up the collector as a service in the OS.
type CreateServiceHandler struct {
	BaseHandler
}

func (h *CreateServiceHandler) Execute(c *CollectorService, selectedCollector *domain.CollectorConfig) error {
	template := c.Templates[selectedCollector.ServiceTemplate]
	err := c.Manager.CreateService(*selectedCollector, selectedCollector.DownloadUrl, template)
	if err != nil {
		return err
	}
	return nil // End of chain, no next handler.
}
