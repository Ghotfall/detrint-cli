package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "detrint-cli",
	Short: "Detrint is like Ansible, only in Golang and with gRPC :)",
	Long:  "Detrint is like Ansible, only in Golang and with gRPC :)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("Failed to start app: %s", err)
	}
}
