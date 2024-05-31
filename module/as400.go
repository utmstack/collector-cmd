package module

import (
	"collector/config.go"
	serv "collector/service"
	"collector/utils"
	"fmt"
	"sync"
)

var (
	as400Collector AS400
	as400Once      sync.Once
)

type AS400 struct {
	Config ProcessConfig
}

func getAS400Collector(host, key string, logger *utils.BeautyLogger) *AS400 {
	as400Once.Do(func() {
		installCmd, installArgs := config.GetAs400Command("install", host, key)
		uninstallCmd, uninstallArgs := config.GetAs400Command("uninstall", "", "")

		as400Collector = AS400{
			Config: ProcessConfig{
				ServiceInfo:           config.GetServiceConfig(config.AS400),
				InstallationCommand:   installCmd,
				InstallationArgs:      installArgs,
				UninstallationCommand: uninstallCmd,
				UninstallationArgs:    uninstallArgs,
				ConnectionHost:        host,
				ConnectionKey:         key,
				Logger:                logger,
			},
		}
	})

	return &as400Collector
}

func (a *AS400) Run() error {
	err := serv.RunService(a.Config.ServiceInfo, a.Config.Logger)
	if err != nil {
		return fmt.Errorf("error running service: %v", err)
	}
	return nil
}

func (a *AS400) Install() error {
	err := a.InstallDependencies()
	if err != nil {
		return fmt.Errorf("error installing dependencies: %v", err)
	}

	result, errB := utils.ExecuteWithResult(a.Config.InstallationCommand, "", a.Config.InstallationArgs...)
	if errB {
		return fmt.Errorf("error executing install command: %v", err)
	}
	err = utils.CheckErrorsInOutput(result)
	if err != nil {
		return fmt.Errorf("error executing install command: %v", err)
	}
	err = serv.InstallService(a.Config.ServiceInfo)
	if err != nil {
		return fmt.Errorf("error installing service: %v", err)
	}

	return nil
}

func (a *AS400) InstallDependencies() error {
	err := InstallJava()
	if err != nil {
		return fmt.Errorf("error installing Java: %v", err)
	}
	return nil
}

func (a *AS400) Uninstall() error {
	result, errB := utils.ExecuteWithResult(a.Config.UninstallationCommand, "", a.Config.UninstallationArgs...)
	if errB {
		return fmt.Errorf("error executing uninstall command: %v", result)
	}

	err := utils.CheckErrorsInOutput(result)
	if err != nil {
		return fmt.Errorf("error executing uninstall command: %v", err)
	}

	err = a.UninstallDependencies()
	if err != nil {
		return fmt.Errorf("error uninstalling dependencies: %v", err)
	}

	err = serv.UninstallService(a.Config.ServiceInfo)
	if err != nil {
		return fmt.Errorf("error uninstalling service: %v", err)
	}

	return nil
}

func (a *AS400) UninstallDependencies() error {
	err := UninstallJava()
	if err != nil {
		return fmt.Errorf("error uninstalling Java: %v", err)
	}
	return nil
}
