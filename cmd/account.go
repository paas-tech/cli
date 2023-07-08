package cmd

import (
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	GroupID: "account",
	Use:     "account",
	Short:   "Get infos about user account",
	Long:    "Get infos about user account",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
