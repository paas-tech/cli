package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "deploy",
	Short:   "Deploy a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deploying project to PaaSTech")

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

		err = project.Deploy(userCfg.GetString("server"), userCfg.GetString("jwt"), projCfg.GetStringMapString("env"))
		if err != nil {
			return err
		}

		fmt.Println("Project successfully deployed")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}
