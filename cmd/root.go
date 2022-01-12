package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type configOptions struct {
	DSN   string
	Debug bool
	Since uint64
	Token string
}

// nolint:gochecknoglobals
var config = &configOptions{
	DSN:   "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
	Debug: false,
}

func newRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cirrus-reporter",
		Short: "bitcoin....",
	}

	cmd.AddCommand(newVersionCmd(version)) // version subcommand
	cmd.AddCommand(NewSetupCommand(), NewBuildsCommand())
	cmd.PersistentFlags().StringVar(&config.DSN, "dsn", config.DSN, "database to connect")
	cmd.PersistentFlags().BoolVarP(&config.Debug, "debug", "d", false, "debug")

	return cmd
}

// Execute invokes the command.
func Execute(version string) error {
	if err := newRootCmd(version).Execute(); err != nil {
		return fmt.Errorf("error executing root command: %w", err)
	}

	return nil
}
