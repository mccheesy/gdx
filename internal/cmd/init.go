package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/mccheesy/gdx/internal/templates"
	"github.com/mccheesy/gdx/internal/tui"
)

// nonAlphanumericRegex is a regular expression to find characters that are not
// lowercase letters, numbers, or hyphens.
var nonAlphanumericRegex = regexp.MustCompile(`[^a-z0-9-]+`)

// slugify converts a string into a URL-friendly "slug".
// It converts the string to lowercase, replaces spaces and special characters
// with hyphens, and removes any trailing hyphens.
func slugify(s string) string {
	slug := strings.ToLower(s)
	slug = nonAlphanumericRegex.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

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

		// Convert the project name to a safe directory name (a "slug").
		dirName := slugify(projectName)
		fmt.Printf("Creating directory: %s\n", dirName)

		// Check if the directory already exists before creating it.
		if _, err := os.Stat(dirName); !os.IsNotExist(err) {
			fmt.Printf("Error: Directory '%s' already exists.\n", dirName)
			os.Exit(1)
		}

		// Create the project directory with standard permissions.
		err := os.Mkdir(dirName, 0755)
		cobra.CheckErr(err)

		// Run the interactive Bubble Tea UI to get user selections.
		choices, err := tui.Run()
		cobra.CheckErr(err)

		if choices == nil {
			fmt.Println("No choices made. Aborting.")
			os.Exit(1)
		}

		// Display the choices to the user
		// In the next step, we will use these to generate files.
		fmt.Println("\nSelected IDEs:")
		fmt.Printf("- %s\n", strings.Join(choices.IDEs, "\n- "))
		fmt.Println("\nSelected Toolset:")
		fmt.Printf("- %s\n", choices.Toolset)

		// Generate the Dockerfile based on the user's choices.
		fmt.Println("Generating Dockerfile...")
		err = templates.GenerateDockerfile(choices, dirName)
		cobra.CheckErr(err)

		// Generate the README.md file.
		fmt.Println("Generating README.md...")
		err = templates.GenerateReadme(choices, projectName, dirName, dirName)
		cobra.CheckErr(err)

		fmt.Println("\n✅ Project created successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
