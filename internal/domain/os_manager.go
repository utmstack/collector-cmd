package domain

type OSManager interface {
	CheckAndInstallDependencies(dependencies []Dependency) error
	DownloadCollector(collector CollectorConfig) (string, error)
	CreateService(config CollectorConfig, servicePath, template string) error
	DisableService(serviceName string) error
}
