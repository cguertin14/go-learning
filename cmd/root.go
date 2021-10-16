package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:           "golearning",
		Short:         "App to extend Go knowledge",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
