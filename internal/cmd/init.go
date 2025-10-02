package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initializes a new Godot project environment",
	Long: `Creates a new project directory, clones a Godot game template,
and generates a custom Dockerfile based on an interactive setup.`,
	Args: cobra.ExactArgs(1), // Ensures exactly one argument (the project name) is passed.
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		fmt.Printf("Initializing project: %s\n", projectName)

		// Placeholder for the directory creation logic.
		// We will implement this in the next step.
		if _, err := os.Stat(projectName); !os.IsNotExist(err) {
			fmt.Printf("Error: Directory '%s' already exists.\n", projectName)
			os.Exit(1)
		}

		fmt.Println("Project directory check passed.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
