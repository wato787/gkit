package main

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var gpCmd = &cobra.Command{
	Use:   "gp [remote] [branch]",
	Short: "Git push command",
	Long:  "Push commits to remote repository",
	Args:  cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkGitRepo(); err != nil {
			return err
		}

		remoteCmd := exec.Command("git", "remote")
		output, err := remoteCmd.Output()
		if err != nil || len(output) == 0 {
			return fmt.Errorf("no remote repository configured")
		}

		if len(args) == 0 {
			return runGitCommand("push")
		}

		gitArgs := append([]string{"push"}, args...)
		return runGitCommand(gitArgs...)
	},
}