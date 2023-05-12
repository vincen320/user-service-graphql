package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vincen320/user-service-graphql/app"
)

var appCmd = &cobra.Command{
	Use:   "run",
	Short: "run the service",
	Run: func(cmd *cobra.Command, args []string) {
		app.Run()
	},
}
