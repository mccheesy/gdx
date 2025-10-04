package templates

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/mccheesy/gdx/internal/tui"
)

// DockerfileData holds all the information needed to generate the Dockerfile.
type DockerfileData struct {
	// IDEs
	Neovide bool
	VSCode  bool
	Rider   bool
	Geany   bool
	// Toolsets
	Is2D bool
	Is3D bool
	// Desktop File Contents
	NeovideDesktopFile string
	VSCodeDesktopFile  string
	GeanyDesktopFile   string
}

// dockerfileTemplate is the Go template for the Dockerfile.
const dockerfileTemplate = `
FROM fedora:latest

RUN adduser --uid 1000 --groups wheel -m gdx && \
    echo "gdx ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

USER gdx

RUN sudo dnf upgrade -y && \
    sudo dnf install -y dotnet-sdk-8.0 git unzip wget zsh && \
    sudo dnf clean all && \
    sudo usermod --shell $(which zsh) gdx

# Install Godot using GodotEnv
RUN dotnet tool install --global Chickensoft.GodotEnv && \
    /home/gdx/.dotnet/tools/godotenv godot install 4.5.0

# Install all selected applications and tools in a single layer to optimize image size
RUN sudo dnf upgrade -y && \
    {{- if .Neovide}}
    sudo dnf copr enable -y agriffis/neovim-nightly && \
    sudo dnf copr enable -y chrisbouchard/neovide-nightly && \
    {{- end}}
    {{- if .VSCode}}
    sudo rpm --import https://packages.microsoft.com/keys/microsoft.asc && \
    sudo sh -c 'echo -e "[code]\nname=Visual Studio Code\nbaseurl=https://packages.microsoft.com/yumrepos/vscode\nenabled=1\ngpgcheck=1\ngpgkey=https://packages.microsoft.com/keys/microsoft.asc" > /etc/yum.repos.d/vscode.repo' && \
    {{- end}}
    # Install all packages at once
    sudo dnf install -y \
    {{- if .Neovide}}
    neovim neovide \
    {{- end}}
    {{- if .VSCode}}
    code \
    {{- end}}
    {{- if .Geany}}
    geany \
    {{- end}}
    {{- if .Is2D}}
    krita tiled \
    {{- end}}
    {{- if .Is3D}}
    blender krita \
    {{- end}}
    # This is a dummy package to ensure dnf install doesn't fail if no options are selected
    which && \
    # Clean up dnf cache to reduce image size
    sudo dnf clean all

{{- if .Neovide}}
# Download Neovide icon
RUN wget https://raw.githubusercontent.com/neovide/neovide/main/assets/neovide.svg -O /usr/share/icons/hicolor/scalable/apps/neovide.svg
RUN cat <<EOF > /usr/share/applications/neovide.desktop
{{.NeovideDesktopFile}}
EOF
{{- end}}

{{- if .VSCode}}
RUN cat <<EOF > /usr/share/applications/code.desktop
{{.VSCodeDesktopFile}}
EOF
{{- end}}

# Come up with a way to install Rider programmatically
# {{- if .Rider}}
# WORKDIR /tmp
#
# RUN wget https://download.jetbrains.com/rider/JetBrains.Rider-2025.2.2.1.tar.gz -O rider.tar.gz && \
#     tar -xzf rider.tar.gz && \
#     # Use a wildcard to handle potential folder name changes between versions
#     mv JetBrains.Rider-* /usr/local/lib/rider && \
#     ln -s /usr/local/lib/rider/bin/rider.sh /usr/local/bin/rider && \
#     rm rider.tar.gz && \
#     wget https://resources.jetbrains.com/storage/products/rider/img/meta/rider_1280x800.svg -O /usr/share/icons/hicolor/scalable/apps/rider.svg
# {{- end}}

# {{- if .Geany}}
# RUN cat <<EOF > /usr/share/applications/geany.desktop
# {{.GeanyDesktopFile}}
# EOF
# {{- end}}

CMD [ "/bin/bash" ]
`

// GenerateDockerfile creates a Dockerfile in the specified project directory
// based on the user's choices.
func GenerateDockerfile(choices *tui.Choices, projectDir string) error {
	data := DockerfileData{
		// Populate desktop file contents from our constants
		NeovideDesktopFile: NeovideDesktop,
		VSCodeDesktopFile:  VSCodeDesktop,
		GeanyDesktopFile:   GeanyDesktop,
	}

	// Populate data from choices
	for _, ide := range choices.IDEs {
		switch ide {
		case "Neovide":
			data.Neovide = true
		case "Visual Studio Code":
			data.VSCode = true
		case "JetBrains Rider":
			data.Rider = true
		case "Geany":
			data.Geany = true
		}
	}

	if strings.HasPrefix(choices.Toolset, "2D") {
		data.Is2D = true
	} else if strings.HasPrefix(choices.Toolset, "3D") {
		data.Is3D = true
	}

	// Create the Dockerfile path
	dockerfilePath := filepath.Join(projectDir, "Dockerfile")
	file, err := os.Create(dockerfilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse and execute the template
	tmpl, err := template.New("dockerfile").Parse(dockerfileTemplate)
	if err != nil {
		return err
	}

	return tmpl.Execute(file, data)
}
