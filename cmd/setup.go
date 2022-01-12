package cmd

import (
	cdb "github.com/pro-bitcoin/cirrus-reporter/pkg/db"
	"github.com/spf13/cobra"
)

func NewSetupCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "setup a new database schema",
		Use:   "setup",
		RunE: func(cmd *cobra.Command, args []string) error {
			db := DBInstance(*config)

			return cdb.Setup(db)
		},
	}

	return cmd
}
