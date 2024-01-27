package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-aiui",
	Short: "A Git interface with chat-based AI",
	Run: func(cmd *cobra.Command, args []string) {
		gitStatus()
	},
}

func gitStatus() {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing Git command:", err)
		return
	}
	fmt.Println(string(output))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
