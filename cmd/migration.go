package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vincen320/user-service-graphql/app"
	"github.com/vincen320/user-service-graphql/migration"
)

func init() {
	// app migration --help || [go run main.go] migration --help
	migrationCmd.PersistentFlags().StringVarP(&migraionFileName, "name", "n", "migration", "name of migration file")
}

var (
	validFirstArgs = map[string]int{
		"up":   1,
		"down": 1,
		"new":  1,
	}
	migraionFileName string
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migrate sql database",
	Args: func(cmd *cobra.Command, args []string) error {
		// Optionally run one of the validators provided by cobra
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		// Run the custom validation logic
		if _, ok := validFirstArgs[args[0]]; !ok {
			return fmt.Errorf("invalid args specified: %s", args[0])
		}
		if len(args) > 1 && (args[0] == "up" || args[0] == "down") {
			if _, err := strconv.Atoi(args[1]); err != nil {
				return fmt.Errorf("invalid args for %s command, should be number", args[0])
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "up", "down":
			tx, err := app.NewDB().Begin()
			if err != nil {
				panic(err)
			}
			limit := 0
			if len(args) > 1 {
				limit, _ = strconv.Atoi(args[1])
			}
			migration := migration.NewMigration(tx)
			err = migration.Migrate(args[0], limit)
			if err != nil {
				tx.Rollback()
			}
			tx.Commit()
		case "new":
			migration := migration.NewMigration(nil)
			migration.New(migraionFileName)
		}
	},
}
