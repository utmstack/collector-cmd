package cmd

import (
	"collector/internal/domain"
	"collector/internal/service"
	"fmt"
	"github.com/spf13/cobra"
)

var host string
var connectionKey string

var enableCmd = &cobra.Command{
	Use:   "enable [collector name] --host [host] --connection-key [key]",
	Short: "Enables a collector service",
	Run: func(cmd *cobra.Command, args []string) {
		collectorName := domain.Collector(args[0])
		osManager := service.NewOSManager()
		collectorService := service.NewCollectorService(osManager, collectorName, host, connectionKey)
		if err := collectorService.SetupCollector(); err != nil {
			fmt.Printf("Failed to enable collector '%s': %v\n", collectorName, err)
		}
		fmt.Printf("%s collector enabled successfully\n", collectorName)
	},
}

var disableCmd = &cobra.Command{
	Use:   "disable [collector name]",
	Short: "Disables a specified collector service",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument: the collector name
	Run: func(cmd *cobra.Command, args []string) {
		collectorName := args[0]
		osManager := service.NewOSManager()
		err := osManager.DisableService(collectorName)
		if err != nil {
			fmt.Printf("Failed to disable collector '%s': %v\n", collectorName, err)
		} else {
			fmt.Printf("Collector '%s' successfully disabled.\n", collectorName)
		}
	},
}

func init() {
	enableCmd.Flags().StringVar(&host, "host", "", "Host for the collector service")
	enableCmd.Flags().StringVar(&connectionKey, "connection-key", "", "Connection key for the collector service")

	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
}
