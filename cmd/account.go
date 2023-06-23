package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/auth"
	"github.com/spf13/cobra"
)

var accountCmd = &cobra.Command{
	GroupID: "account",
	Use:     "account",
	Short:   "Get infos about user account",
	Long:    "Get infos about user account",
	RunE: func(cmd *cobra.Command, args []string) error {
		status, err := auth.Status()
		fmt.Println(status)
		return err
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)
}
