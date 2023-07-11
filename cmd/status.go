package cmd

import (
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "status",
	Short:   "List deployments for current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
