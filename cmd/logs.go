package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	GroupID: "deployment",
	Use:     "logs",
	Short:   "Get logs from a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deployment logs:")

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

		logs, err := project.Logs(userCfg.GetString("server"), userCfg.GetString("jwt"))
		if err != nil {
			return err
		}

		fmt.Print(logs)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
