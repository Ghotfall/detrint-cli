package main

import (
	"github.com/ghotfall/detrint-cli/cmd"
	"go.uber.org/zap"
	"log"
)

func main() {
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

	cmd.Execute(logger)
}
