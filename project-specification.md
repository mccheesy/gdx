Project Specification: gdx
1. Project Title

gdx: A Godot Development eXperience CLI
2. Overview

gdx is a command-line interface (CLI) tool that manages project-specific, containerized development environments for Godot. It simplifies the setup process and provides a seamless workflow on container-focused desktops like Bazzite and Universal Blue.

The primary command, gdx init, bootstraps a new project by:

    Cloning a base Godot project template.

    Generating a tailored Dockerfile that includes the user's choice of IDE and a suite of tools for either 2D or 3D development.

    Creating a README.md with personalized instructions on how to build the container and get started.

Future commands will allow users to manage and interact with the project's environment directly from the command line.
3. Core Features (for the init command)

    Interactive Setup: The init command will guide the user through a series of simple questions to configure their environment.

    Project Template: It will initialize the project by cloning the Chickensoft GodotGame repository as a starting point.

    IDE Selection: Users can choose one or more IDEs to be included in the container.

        Neovide (via Neovim Nightly)

        Visual Studio Code

        JetBrains Rider

        Geany

    Toolset Selection: Users can select a pre-configured suite of popular open-source tools tailored for a specific workflow.

        2D Toolset: Aseprite, Krita, Tiled Map Editor.

        3D Toolset: Blender, Krita, Aseprite.

    Dynamic Dockerfile Generation: The script will construct a Dockerfile on the fly based on user selections.

    Documentation Generation: A README.md will be generated with instructions specific to the created environment.

4. Command-Line Interface (CLI) Design

gdx will be a standalone CLI with multiple subcommands.
The init Command

This command bootstraps a new project.

gdx init "My Awesome Game"

Upon execution, it will create the project directory, clone the template, and initiate the interactive setup prompts as previously defined.
Future Command Structure (Examples)

The CLI will be extensible to include commands for managing the environment after initialization. This provides a single, consistent interface for all project-related tasks.

# (In the project directory)

# Build the project's container image
gdx build

# Launch an application from the project's container
gdx launch godot
gdx launch aseprite

# Enter the project's container shell
gdx enter

5. Technical Implementation Details

    Language: The CLI should be written in Go (Golang) to produce a single, statically-linked binary that is easily distributable across Linux, Windows, and macOS.

    Recommended Go Libraries:

        Command Structure: Cobra is the standard for creating powerful, multi-command CLI applications.

        Interactive Prompts: The Bubble Tea framework is an excellent choice for building polished, interactive terminal UIs.

    Dockerfile Construction: The tool will build the Dockerfile by concatenating modular text blocks based on user input.

    Package Management: All tool installations within the generated Dockerfile should use Fedora's dnf package manager.

    Configuration Files: The tool will generate the necessary .desktop files for the chosen applications and place them in a desktop-files subdirectory.

6. Generated Project Structure

The output of the gdx init command remains the same, creating a complete, self-contained project directory.

my-awesome-game/
├── .git/
├── src/
├── project.godot
├── Dockerfile
├── README.md
├── vscode.repo         # (If VS Code is selected)
└── desktop-files/
    ├── godot-mono.desktop
    └── ... (other selected app .desktop files)

7. Example Usage Flow

    User runs gdx init "My Kenshi Clone".

    Script creates the directory and clones the template.

    Script prompts the user, who selects "JetBrains Rider" and the "3D" toolset.

    The script generates a Dockerfile and all necessary configuration files.

    The script outputs a final success message:
    ✅ Project "My Kenshi Clone" created successfully!
       To get started, run the following commands:
       cd "My Kenshi Clone"
       gdx build
