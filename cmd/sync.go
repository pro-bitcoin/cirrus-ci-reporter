package cmd

import "github.com/spf13/cobra"

func NewSyncCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "sync builds and artifacts",
		RunE: func(c *cobra.Command, args []string) error {
			return getBuilds(c.Context(), *config)
		},
	}
	cmd.AddCommand(NewBuildsCommand())
	return cmd
}
