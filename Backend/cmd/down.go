package cmd

import (
	"example/Project3/internal/migration"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			panic("length of arguments cannot be more than 1")
		}
		// fmt.Println(args)
		m := &migration.Migrate{}

		err := m.Down(args)
		if err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	migrateCmd.AddCommand(downCmd)
}
