package cmd

import (
	"fmt"

	"github.com/paastech-cloud/cli/pkg/deployment"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "deploy",
	Short:   "Deploy a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Deploying project")
		return deployment.Deploy()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
