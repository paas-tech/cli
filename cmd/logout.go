package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	GroupID: "account",
	Use:     "logout",
	Short:   "Log out from PaaSTech.cloud",
	Long:    "Log out from PaaSTech.cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load auth config
		err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		// Save empty jwt in auth conf
		config.SetAuth("", "")
		fmt.Println("🚪 Logging out")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
