# gdx: A Godot Development Environment CLI

**`gdx`** is a command-line interface (CLI) tool that manages project-specific, containerized development environments for Godot. It simplifies the setup process and provides a seamless workflow on container-focused desktops like Bazzite and Universal Blue.

The primary command, `gdx init`, bootstraps a new project by:

*   Cloning a base Godot project template.
*   Generating a tailored `Dockerfile` that includes the user's choice of IDE and a suite of tools for either 2D or 3D development.
*   Creating a `README.md` with personalized instructions on how to build the container and get started.

This project is being developed with assistance from Gemini Code Assist. See `GEMINI.md` for our shared project context and development plan.

## Development Setup

To contribute to `gdx`, you'll need to set up a Go development environment.

### Prerequisites

*   **Go**: Version 1.21 or later.

### Installation

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/mccheesy/gdx.git
    cd gdx
    ```

2.  **Install Dependencies:**

    This project uses Go Modules to manage dependencies. They will be automatically downloaded when you build or run the project.

    ```sh
    go mod tidy
    ```

3.  **Build and Run:**

    You can run the application directly using `go run` or build a binary.

    ```sh
    # Run the application from the project root
    go run ./cmd/gdx --help

    # Or, build the binary
    go build ./cmd/gdx
    ./gdx --help
    ```

## Usage

The primary command is `init`. It takes one argument: the name of your project.

```sh
go run ./cmd/gdx init "My Awesome Game"
```
