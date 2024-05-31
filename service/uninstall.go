package serv

import (
	"collector/config.go"
	"collector/utils"
	"fmt"
)

func UninstallService(cnf config.ServiceConfig) error {
	utils.StopService(cnf.Name)
	err := utils.UninstallService(cnf.Name)
	if err != nil {
		return fmt.Errorf("error uninstalling service: %v", err)
	}
	return nil
}
