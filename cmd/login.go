package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	username string
	password string
)

var loginCmd = &cobra.Command{
	GroupID: "account",
	Use:     "login",
	Short:   "Log in to PaaSTech.cloud",
	Long:    "Log in to PaaSTech.cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		if len(username) == 0 || len(password) == 0 {
			fmt.Print("Username: ")
			fmt.Scan(&username)

			fmt.Print("Password: ")
			p, _ := terminal.ReadPassword(0)
			password = string(p)
			fmt.Print("\n")
		}

		fmt.Printf("🔐 Logging in as %s...\n", username)
		err := auth.Login(username, password)
		if err == nil {
			fmt.Println("✅ Login successful")
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagsRequiredTogether("username", "password")
}
