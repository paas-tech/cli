package cmd

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
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

		// Check if current directory is a git repo
		_, err := git.PlainOpen(".")
		if err != nil {
			return errors.New("No git repository found in current directory")
		}

		// Load User config
		userCfg, err := config.LoadAuthConfig()
		if err != nil {
			return err
		}

		// Check if user is authenticated
		connected, err := config.IsAuthenticated(userCfg)
		if err != nil {
			return err
		}
		if !connected {
			return errors.New("Not logged in")
		}

		// Check if project exists
		if config.ProjectExists() {
			return errors.New("Project already exists")
		}

		// Create project
		project, err := project.CreateProject(userCfg.GetString("server"), userCfg.GetString("jwt"), args[0])
		if err != nil {
			return err
		}

		// Create project conf
		err = config.CreateProjectConfig()
		if err != nil {
			return err
		}

		// Load project config
		projCfg, err := config.LoadProjectConfig()
		if err != nil {
			return err
		}

		// Write config
		projCfg.Set("project", project)
		projCfg.WriteConfig()

		fmt.Println("Project " + project.Name + " created.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
