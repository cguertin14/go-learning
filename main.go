package main

import (
	"context"

	"github.com/cguertin14/go-learning/cmd"
	"github.com/cguertin14/logger"
)

func main() {
	// Context setup
	ctx := context.Background()
	ctxLogger := logger.Initialize(logger.Config{
		Level: "debug",
	})
	ctx = context.WithValue(ctx, logger.CtxKey, ctxLogger)

	// Run command
	if err := cmd.Execute(ctx); err != nil {
		ctxLogger.Fatalf("error when processing command: %s", err.Error())
	}
}
