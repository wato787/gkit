package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func isGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()
	return err == nil
}

func checkGitRepo() error {
	if !isGitRepo() {
		return fmt.Errorf("not a git repository")
	}
	return nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

var rootCmd = &cobra.Command{
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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}