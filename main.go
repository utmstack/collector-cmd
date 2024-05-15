package main

import (
	"collector/config.go"
	"collector/module"
	"collector/utils"
	"fmt"
	"os"
)

func main() {
	beautyLogger := utils.GetBeautyLogger()
	beautyLogger.PrintBanner()

	if len(os.Args) > 1 {
		if isValid := config.ValidateParams(os.Args); !isValid {
			beautyLogger.WriteFatal("Invalid parameters", fmt.Errorf("invalid parameters"))
		}

		action := os.Args[1]
		typ := config.Collector(os.Args[2])
		host := ""
		connectionKey := ""

		if len(os.Args) > 4 {
			host = os.Args[3]
			connectionKey = os.Args[4]
		}

		collector := module.GetCollectorProcess(typ, host, connectionKey)

		switch action {
		case "run":
			err := collector.Run()
			if err != nil {
				beautyLogger.WriteFatal(fmt.Sprintf("Error running %s collector:", typ), err)
			}
		case "install":
			beautyLogger.WriteSimpleMessage(fmt.Sprintf("Installing %s collector", typ))
			config.SaveConfig(&config.ServiceTypeConfig{CollectorType: typ})
			err := collector.Install()
			if err != nil {
				beautyLogger.WriteFatal(fmt.Sprintf("Error installing %s collector:", typ), err)
			} else {
				beautyLogger.WriteSimpleMessage(fmt.Sprintf("%s collector installed successfully", typ))
			}
		case "uninstall":
			beautyLogger.WriteSimpleMessage(fmt.Sprintf("Uninstalling %s collector", typ))
			err := collector.Uninstall()
			if err != nil {
				beautyLogger.WriteFatal(fmt.Sprintf("Error uninstalling %s collector:", typ), err)
			} else {
				beautyLogger.WriteSimpleMessage(fmt.Sprintf("%s collector uninstalled successfully", typ))
			}
		}
	} else {
		cnf, err := config.ReadConfig()
		if err != nil {
			beautyLogger.WriteFatal("Error reading config", err)
		}
		collector := module.GetCollectorProcess(config.Collector(cnf.CollectorType), "", "")
		err = collector.Run()
		if err != nil {
			beautyLogger.WriteFatal(fmt.Sprintf("Error running %s collector:", cnf.CollectorType), err)
		}
	}
}
