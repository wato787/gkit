package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gkit-installer",
	Short: "Install all gkit commands",
	Long:  "Install gs, ga, gc, and gp commands from gkit package",
	RunE: func(cmd *cobra.Command, args []string) error {
		return installCommands()
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func installCommands() error {
	commands := []string{"gs", "ga", "gc", "gp"}
	
	fmt.Println("Installing gkit commands...")
	
	for _, command := range commands {
		fmt.Printf("Installing %s...", command)
		
		cmd := exec.Command("go", "install", fmt.Sprintf("github.com/wato787/gkit/cmd/%s@latest", command))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		
		if err := cmd.Run(); err != nil {
			fmt.Printf(" ❌ Failed\n")
			return fmt.Errorf("failed to install %s: %v", command, err)
		}
		
		fmt.Printf(" ✅ Installed\n")
	}
	
	fmt.Println("\n🎉 All gkit commands installed successfully!")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  gs - git switch")
	fmt.Println("  ga - git add") 
	fmt.Println("  gc - git commit")
	fmt.Println("  gp - git push")
	
	// Check if GOBIN or GOPATH/bin is in PATH
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			// Default GOPATH
			home, _ := os.UserHomeDir()
			gopath = filepath.Join(home, "go")
		}
		gobin = filepath.Join(gopath, "bin")
	}
	
	path := os.Getenv("PATH")
	if !contains(path, gobin) {
		fmt.Printf("\n⚠️  Warning: %s is not in your PATH\n", gobin)
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
			fmt.Printf("Add this to your shell profile (~/.bashrc, ~/.zshrc, etc.):\n")
			fmt.Printf("export PATH=\"%s:$PATH\"\n", gobin)
		}
	}
	
	return nil
}

func contains(path, dir string) bool {
	pathDirs := strings.Split(path, ":")
	for _, pathDir := range pathDirs {
		if pathDir == dir {
			return true
		}
	}
	return false
}