package main

import (
	"fmt"
	"os"
	"os/exec"

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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}