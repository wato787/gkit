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

var rootCmd = &cobra.Command{
	Use:   "gs [branch]",
	Short: "Git switch command",
	Long:  "Switch to a branch or create a new branch with prefix expansion",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := checkGitRepo(); err != nil {
			return err
		}

		createBranch, _ := cmd.Flags().GetBool("create")

		if len(args) == 0 {
			if createBranch {
				return fmt.Errorf("branch name required when using -c flag")
			}
			return runGitCommand("switch")
		}

		branch := args[0]

		if branch == "-" {
			return runGitCommand("switch", "-")
		}

		if createBranch {
			expandedBranch := expandBranchPrefix(branch)
			
			cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+expandedBranch)
			if cmd.Run() == nil {
				return fmt.Errorf("branch '%s' already exists", expandedBranch)
			}
			
			return runGitCommand("switch", "-c", expandedBranch)
		}

		checkCmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branch)
		if checkCmd.Run() != nil {
			return fmt.Errorf("branch '%s' does not exist", branch)
		}

		return runGitCommand("switch", branch)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("create", "c", false, "Create a new branch")
	
	rootCmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) == 0 {
			createFlag, _ := cmd.Flags().GetBool("create")
			if createFlag {
				prefixes := []string{
					"feature/",
					"fix/",
					"hotfix/",
					"release/",
					"bugfix/",
					"epic/",
				}
				
				if toComplete == "f" {
					return []string{"feature/", "fix/"}, cobra.ShellCompDirectiveNoSpace
				}
				
				var completions []string
				for _, prefix := range prefixes {
					if strings.HasPrefix(prefix, toComplete) {
						completions = append(completions, prefix)
					}
				}
				return completions, cobra.ShellCompDirectiveNoSpace
			}
			
			cmd := exec.Command("git", "branch", "--format=%(refname:short)")
			output, err := cmd.Output()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			
			branches := strings.Split(strings.TrimSpace(string(output)), "\n")
			var completions []string
			for _, branch := range branches {
				if strings.HasPrefix(branch, toComplete) {
					completions = append(completions, branch)
				}
			}
			return completions, cobra.ShellCompDirectiveDefault
		}
		return nil, cobra.ShellCompDirectiveDefault
	}
}