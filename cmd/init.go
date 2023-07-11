package cmd

import (
	"errors"
	"fmt"

	"github.com/paastech-cloud/cli/internal/config"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: "project",
	Use:     "init [Project Name]",
	Short:   "Initialize a project",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Initializing a new project")

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

		if config.ProjectExists() {
			return errors.New("Project already exists")
		}

		project, err := project.CreateProject(userCfg.GetString("server"), userCfg.GetString("jwt"), args[0])
		if err != nil {
			return err
		}

		err = config.CreateProjectConfig()
		if err != nil {
			return err
		}

		projCfg, err := config.LoadProjectConfig()
		if err != nil {
			return err
		}

		projCfg.Set("project", project)
		projCfg.WriteConfig()

		fmt.Println("Project " + project.Name + " created.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
