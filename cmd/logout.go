package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/auth"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	GroupID: "account",
	Use:     "logout",
	Short:   "Log out from PaaSTech.cloud",
	Long:    "Log out from PaaSTech.cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("🚪 Logging out")
		return auth.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
