package config

import (
	"collector/internal/domain"
	"encoding/json"
	"os"
	"sync"
)

var (
	once               sync.Once
	collectorConfig    *[]domain.CollectorConfig
	serviceTemplateMap map[string]string
)

func LoadCollectorConfig(path string) *[]domain.CollectorConfig {
	once.Do(func() {
		file, err := os.ReadFile(path)
		if err != nil {
			panic("Unable to read collector config file: " + err.Error())
		}
		err = json.Unmarshal(file, &collectorConfig)
		if err != nil {
			panic("Unable to parse collector config: " + err.Error())
		}
	})
	return collectorConfig
}

func LoadServiceTemplates(path string) map[string]string {
	once.Do(func() {
		file, err := os.ReadFile(path)
		if err != nil {
			panic("Unable to read service template file: " + err.Error())
		}
		err = json.Unmarshal(file, &serviceTemplateMap)
		if err != nil {
			panic("Unable to parse service template config: " + err.Error())
		}
	})
	return serviceTemplateMap
}
