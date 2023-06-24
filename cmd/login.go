package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	email    string
	password string
)

var loginCmd = &cobra.Command{
	GroupID: "account",
	Use:     "login",
	Short:   "Log in to PaaSTech.cloud",
	Long:    "Log in to PaaSTech.cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		if len(email) == 0 || len(password) == 0 {
			fmt.Print("Email: ")
			fmt.Scan(&email)

			fmt.Print("Password: ")
			p, _ := terminal.ReadPassword(0)
			password = string(p)
			fmt.Print("\n")
		}

		fmt.Printf("🔐 Logging in as %s...\n", email)
		err := auth.Login(email, password)
		if err == nil {
			fmt.Println("✅ Login successful")
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagsRequiredTogether("email", "password")
}
