package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "detrint-cli",
	Short: "Detrint is like Ansible, only in Golang and with gRPC :)",
	Long: `Detrint is a configuration management tool developed using Golang & gRPC.
The main focus of the tool is speed and usability for DevOps Engineers`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Failed to start app: %v", err)
	}
}
