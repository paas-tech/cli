package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	GroupID: "account",
	Use:     "logout",
	Short:   "Log out from PaaSTech",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Load auth config
		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		// Save empty jwt in auth conf
		userCfg.Set("server", "")
		userCfg.Set("jwt", "")
		userCfg.WriteConfig()

		fmt.Println("🚪 Logging out")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
