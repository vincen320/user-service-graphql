package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(migrationCmd, appCmd)
}

var rootCmd = &cobra.Command{Use: "app"}

func Execute() error {
	return rootCmd.Execute()
}
