package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var gcCmd = &cobra.Command{
	Use:   "gc [message]",
	Short: "Git commit command",
	Long:  "Commit staged changes with a message",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkGitRepo(); err != nil {
			return err
		}

		statusCmd := exec.Command("git", "diff", "--cached", "--name-only")
		output, err := statusCmd.Output()
		if err != nil {
			return fmt.Errorf("failed to check staged files: %v", err)
		}

		if strings.TrimSpace(string(output)) == "" {
			return fmt.Errorf("no changes added to commit")
		}

		if len(args) == 0 {
			return runGitCommand("commit")
		}

		return runGitCommand("commit", "-m", args[0])
	},
}