package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paastech",
	Short: "PaaSTech CLI is a tool to manage your PaaSTech applications",
	Long: `
🍉 PaaSTech CLI is the official tool to manage and operate your 
   applications and deployments on the PaaSTech platform.

🔗 Find more information at: https://paastech.cloud/
	`,
	SilenceUsage: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddGroup(&cobra.Group{ID: "account", Title: "Account Commands:"})
	rootCmd.AddGroup(&cobra.Group{ID: "project", Title: "Project Commands:"})
}
