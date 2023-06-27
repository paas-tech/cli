package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/paastech-cloud/cli/pkg/project"
	"github.com/paastech-cloud/cli/pkg/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	GroupID: "project",
	Use:     "delete",
	Short:   "Delete a project",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting this project from PaasTech")

		// check if git repo exists
		git, err := git.PlainOpen(".")
		if err != nil {
			return errors.New("No git repository found in current directory")
		}

		// confirmation
		stdin := bufio.NewReader(os.Stdin)
		if utils.ConfirmationPrompt(stdin) {
			return project.DeleteProject(git)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
