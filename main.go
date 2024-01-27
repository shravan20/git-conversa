package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-conversa",
	Short: "A Git interface with chat-based AI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Git Conversa! Use 'git-conversa(gc-tools) help' to see available commands.")
	},
}

// Command definitions
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add files to the staging area",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("add", args)
	},
}

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create, list, or delete branches",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("branch", args)
	},
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Switch branches",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("checkout", args)
	},
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit changes to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("commit", args)
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download objects from a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("fetch", args)
	},
}

var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "Search for text in the repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("grep", args)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("init", args)
	},
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the commit history",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("log", args)
	},
}

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge branches",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("merge", args)
	},
}

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch and merge changes from a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("pull", args)
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to a remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("push", args)
	},
}

var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Apply commits from one branch to another",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("rebase", args)
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Unstage or reset changes",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("reset", args)
	},
}

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files from the repository",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("rm", args)
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information about a commit, tag, or branch",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("show", args)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the working directory and staging area",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("status", args)
	},
}

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Create, list, or delete tags",
	Run: func(cmd *cobra.Command, args []string) {
		gitCommand("tag", args)
	},
}

// Function to execute Git commands
func gitCommand(command string, args []string) {
	cmd := exec.Command("git", append([]string{command}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func main() {
	// Add commands to the root command
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(branchCmd)
	rootCmd.AddCommand(checkoutCmd)
	rootCmd.AddCommand(commitCmd)
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(grepCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(mergeCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(rebaseCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(rmCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(tagCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
