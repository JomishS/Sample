package cmd

import (
	"example/Project3/internal/server"
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Init()
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
