package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/auth"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	email    string
	password string
	server   string
)

var loginCmd = &cobra.Command{
	GroupID: "account",
	Use:     "login",
	Short:   "Log in to PaaSTech",
	RunE: func(cmd *cobra.Command, args []string) error {
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		// Ask for email and password if not present in flags
		if len(email) == 0 || len(password) == 0 {
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			fmt.Print("Password: ")
			p, _ := terminal.ReadPassword(0)
			password = string(p)
			fmt.Print("\n")
		}
		if len(email) == 0 || len(password) == 0 {
			return errors.New("Email and password cannot be empty. Please try again.")
		}

		// Use PaaSTech API server by default
		if server == "" {
			server = "https://api.paastech.cloud"
		} else {
			server = strings.TrimSuffix(server, "/")
		}

		// Send login request
		fmt.Printf("🔐 Logging in as %s\n", email)
		jwt, err := auth.Login(server, email, password)
		if err != nil {
			return err
		}

		// Load auth config
		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		// Save jwt in auth conf
		userCfg.Set("server", server)
		userCfg.Set("jwt", jwt)
		userCfg.WriteConfig()

		fmt.Println("✅ Login successful")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&server, "server", "", "", "Server URL")
	loginCmd.Flags().StringVarP(&email, "email", "e", "", "Email")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	loginCmd.MarkFlagsRequiredTogether("email", "password")
}
