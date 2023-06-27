package cmd

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	GroupID: "project",
	Use:     "init",
	Short:   "Initialize a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Initializing a new project")

		git, err := git.PlainOpen(".")
		if err != nil {
			return errors.New("no git repository found in current directory")
		}

		return project.InitProject(git)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
