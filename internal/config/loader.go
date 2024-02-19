package config

import (
	"collector/internal/domain"
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

var (
	once               sync.Once
	collectorConfig    *[]domain.CollectorConfig
	serviceTemplateMap map[string]string
)

func LoadCollectorConfig() *[]domain.CollectorConfig {
	once.Do(func() {
		path, _ := getExecPath()
		path = filepath.Join(path, "templates/collectors.json")
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

func LoadServiceTemplates() map[string]string {
	once.Do(func() {
		path, _ := getExecPath()
		path = filepath.Join(path, "templates/service_tmpl.json")
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

func getExecPath() (string, error) {
	// Get the path to the executable.
	//execPath, err := os.Executable()
	//if err != nil {
	//	return "", fmt.Errorf("failed to get executable path: %w", err)
	//}
	//
	//// Get the directory of the executable.
	//execDir := filepath.Dir(execPath)

	// Construct the path to the configuration file
	return "/Users/jd/Documents/WORKSPACE/UTM/collector-cmd", nil
	//return execDir, nil
}
