package cmd

import (
	"github.com/godazz/bloGin/pkg/bootstrap"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "database seeder",
	Long:  "database seeder",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}

func seed() {
	bootstrap.Seed()
}
