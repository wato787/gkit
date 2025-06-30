package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gkit",
	Short: "Git CLI tools collection",
	Long:  "A collection of Git CLI tools to simplify common git operations",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(gsCmd)
	rootCmd.AddCommand(gaCmd)
	rootCmd.AddCommand(gcCmd)
	rootCmd.AddCommand(gpCmd)
}