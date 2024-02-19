package domain

type Collector string

var (
	as400 Collector = "as400"
)

type Dependency struct {
	Name           Collector         `json:"name"`
	CheckCommand   string            `json:"checkCommand"`
	InstallCommand map[string]string `json:"installCommand"`
}

type CollectorConfig struct {
	Name            Collector    `json:"name"`
	Type            string       `json:"type"`
	DownloadUrl     string       `json:"downloadUrl"`
	Dependencies    []Dependency `json:"dependencies"`
	ServiceTemplate string       `json:"serviceTemplate"`
}
