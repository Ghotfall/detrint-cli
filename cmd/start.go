package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start deployment",
	Long:  "Start deployment using provided state file and inventory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(statePath)
		fmt.Println(inventoryPath)
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
	startCmd.MarkFlagRequired("state")

	startCmd.Flags().StringVarP(
		&inventoryPath,
		"inventory",
		"i",
		"",
		"Path to inventory file (required)",
	)
	startCmd.MarkFlagRequired("inventory")

	rootCmd.AddCommand(startCmd)
}
