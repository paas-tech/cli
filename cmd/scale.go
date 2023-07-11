package cmd

import (
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "scale [number]",
	Short:   "Scale a project deployment",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scaleCmd)
}
