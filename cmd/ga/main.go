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

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}