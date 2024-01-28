package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// GitConversa
var rootCmd = &cobra.Command{
	Use:   "git-conversa",
	Short: "A Git interface with chat-based AI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Git Conversa! Use 'git-conversa help' to see available commands.")
		p := prompt.New(
			func(t string) { // onInput
				t = strings.TrimSpace(t)
				if t == "" {
					return
				}

				if t == "exit" || t == "quit" {
					os.Exit(0)
				}

				parts := strings.Fields(t)
				if len(parts) > 0 {
					command := parts[0]
					args := parts[1:]
					executeCommand(command, args)
				}
			},
			completer,
			prompt.OptionPrefix("> "),
			prompt.OptionTitle("Git Conversa"),
			prompt.OptionInputTextColor(prompt.Yellow),
		)

		p.Run()
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
		message, _ := cmd.Flags().GetString("message")
		if message == "" {
			// Prompt for message if not provided
			fmt.Print("Enter commit message: ")
			fmt.Scanln(&message)
		}

		author, _ := cmd.Flags().GetString("author")
		if author == "" {
			// Prompt for author if not provided
			fmt.Print("Enter commit author: ")
			fmt.Scanln(&author)
		}

		// Validate and use the input values
		if message == "" {
			fmt.Println("Error: Commit message cannot be empty.")
			os.Exit(1)
		}

		gitCommand("commit", []string{"-m", message, "--author", author})
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

// Function to execute Git commands
func executeCommand(command string, args []string) {
	cmd := exec.Command("git", append([]string{command}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "add", Description: "Add files to the staging area"},
		{Text: "branch", Description: "Create, list, or delete branches"},
		{Text: "checkout", Description: "Switch branches"},
		{Text: "commit", Description: "Commit changes to the repository"},
		{Text: "fetch", Description: "Download objects from a remote repository"},
		{Text: "grep", Description: "Search for text in the repository"},
		{Text: "init", Description: "Initialize a new repository"},
		{Text: "log", Description: "Show the commit history"},
		{Text: "merge", Description: "Merge branches"},
		{Text: "pull", Description: "Fetch and merge changes from a remote repository"},
		{Text: "push", Description: "Push changes to a remote repository"},
		{Text: "rebase", Description: "Apply commits from one branch to another"},
		{Text: "reset", Description: "Unstage or reset changes"},
		{Text: "rm", Description: "Remove files from the repository"},
		{Text: "show", Description: "Show information about a commit, tag, or branch"},
		{Text: "status", Description: "Show the status of the working directory and staging area"},
		{Text: "tag", Description: "Create, list, or delete tags"},
	}
	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
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

	// Add flags to the commit command
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	commitCmd.Flags().String("author", "", "Commit author")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
