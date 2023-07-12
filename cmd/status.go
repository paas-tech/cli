package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "status",
	Short:   "Get status of a deployment",
	RunE: func(cmd *cobra.Command, args []string) error {
		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		connected, err := config.IsAuthenticated(userCfg)
		if err != nil {
			return err
		}
		if !connected {
			return errors.New("Not logged in")
		}

		projCfg, err := config.LoadProjectConfig()
		if err != nil {
			return err
		}

		var project project.Project
		projCfg.UnmarshalKey("project", &project)

		status, err := project.Status(userCfg.GetString("server"), userCfg.GetString("jwt"))
		if err != nil {
			return err
		}

		fmt.Print("Deployment status: " + status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
