package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"log"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start deployment",
	Long:  "Start deployment using provided state file and inventory",
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to create logger: %s", err)
		}
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {
				logger.Debug("Failed to sync logger", zap.Error(err))
			}
		}(logger)

		logger.Info("Start is called!")
	},
}

var statePath string
var inventoryPath string

func init() {
	startCmd.Flags().StringVarP(
		&statePath,
		"state",
		"s",
		"",
		"Path to state file (required)",
	)
	err := startCmd.MarkFlagRequired("state")
	if err != nil {
		fmt.Printf("An error occured: %s", err)
	}

	startCmd.Flags().StringVarP(
		&inventoryPath,
		"inventory",
		"i",
		"",
		"Path to inventory file (required)",
	)
	err = startCmd.MarkFlagRequired("inventory")
	if err != nil {
		fmt.Printf("An error occured: %s", err)
	}

	rootCmd.AddCommand(startCmd)
}
