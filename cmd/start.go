package cmd

import (
	"fmt"
	"github.com/ghotfall/detrint/inv"
	"github.com/ghotfall/detrint/state"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io/ioutil"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start deployment",
	Long:  "Start deployment using provided state file and inventory",
	Run: func(cmd *cobra.Command, args []string) {
		stateSet, err := loadState(statePath)
		if err != nil {
			fmt.Printf("Failed to read state file: %v", err)
			return
		}

		inventory, err := loadInventory(inventoryPath)
		if err != nil {
			fmt.Printf("Failed to read inventory file: %v", err)
			return
		}

		logger, err := zap.NewProduction()
		if err != nil {
			fmt.Printf("Failed to create logger: %v", err)
			return
		}
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {
				logger.Debug("Failed to sync logger", zap.Error(err))
			}
		}(logger)

		stateSet.Deploy(*inventory, logger)
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
		fmt.Printf("An error occured: %v", err)
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
		fmt.Printf("An error occured: %v", err)
	}

	rootCmd.AddCommand(startCmd)
}

func loadInventory(filename string) (*inv.Inventory, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %v", err)
	}

	var i inv.Inventory
	err = toml.Unmarshal(file, &i)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %v", err)
	}
	return &i, nil
}

func loadState(filename string) (*state.Set, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file data: %v", err)
	}

	var s state.Set
	err = toml.Unmarshal(file, &s)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %v", err)
	}
	return &s, nil
}
