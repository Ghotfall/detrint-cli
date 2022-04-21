package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger *zap.Logger

var rootCmd = &cobra.Command{
	Use:   "detrint-cli",
	Short: "Detrint is like Ansible, only in Golang and with gRPC :)",
	Long:  "Detrint is like Ansible, only in Golang and with gRPC :)",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("App is started!")
	},
}

func Execute(l *zap.Logger) {
	logger = l

	err := rootCmd.Execute()
	if err != nil {
		logger.Fatal("Failed to start app", zap.Error(err))
	}
}
