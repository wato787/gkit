package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var gsCmd = &cobra.Command{
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

func init() {
	gsCmd.Flags().BoolP("create", "c", false, "Create a new branch")
	
	gsCmd.ValidArgsFunction = func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
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