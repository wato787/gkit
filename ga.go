package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var gaCmd = &cobra.Command{
	Use:   "ga [files...]",
	Short: "Git add command",
	Long:  "Add files to the staging area",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkGitRepo(); err != nil {
			return err
		}

		for _, arg := range args {
			if arg != "." && arg != "*" {
				if _, err := os.Stat(arg); os.IsNotExist(err) {
					return fmt.Errorf("pathspec '%s' did not match any files", arg)
				}
			}
		}

		gitArgs := append([]string{"add"}, args...)
		return runGitCommand(gitArgs...)
	},
}