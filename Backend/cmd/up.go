package cmd

import (
	"example/Project3/internal/migration"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Migrate up command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			panic("length of arguments cannot be more than 1")
		}
		m := &migration.Migrate{}

		err := m.Up(args)
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	migrateCmd.AddCommand(upCmd)
}
