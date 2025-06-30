package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

func expandBranchPrefix(branch string) string {
	prefixMap := map[string]string{
		"f/": "feature/",
		"fix/": "fix/",
		"h/": "hotfix/",
		"r/": "release/",
		"b/": "bugfix/",
		"e/": "epic/",
	}
	
	for prefix, expansion := range prefixMap {
		if strings.HasPrefix(branch, prefix) {
			return strings.Replace(branch, prefix, expansion, 1)
		}
	}
	return branch
}