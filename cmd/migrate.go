package cmd

import (
	"github.com/godazz/bloGin/pkg/bootstrap"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Table migration",
	Long:  "A command to migrate database tables",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate() {
	bootstrap.Migrate()
}
