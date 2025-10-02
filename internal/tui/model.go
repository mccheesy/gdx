package tui

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Choices holds the user's final selections from the TUI.
type Choices struct {
	IDEs    []string
	Toolset string
}

// step represents which part of the interactive setup we are in.
type step int

const (
	stepIDE step = iota
	stepToolset
)

type model struct {
	step           step
	ideOptions     []string        // All available IDEs
	selectedIDEs   map[string]bool // The IDEs the user has selected
	toolsetOptions []string        // All available toolsets
	cursor         int             // Which item our cursor is pointing at
	quitting       bool            // Flag to indicate the TUI should exit
	finalChoices   *Choices        // The final selections to be returned
}

func initialModel() model {
	return model{
		step:           stepIDE,
		ideOptions:     []string{"Neovide", "Visual Studio Code", "JetBrains Rider", "Geany"},
		selectedIDEs:   make(map[string]bool),
		toolsetOptions: []string{"2D (Aseprite, Krita, Tiled)", "3D (Blender, Krita, Aseprite)"},
		cursor:         0,
	}
}

// Init is the first function that will be called. It returns a command.
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			optionsLen := 0
			if m.step == stepIDE {
				optionsLen = len(m.ideOptions)
			} else {
				optionsLen = len(m.toolsetOptions)
			}
			if m.cursor < optionsLen-1 {
				m.cursor++
			}

		case "enter":
			if m.step == stepIDE {
				// Move to the next step
				m.step = stepToolset
				m.cursor = 0
			} else {
				// Final step, so we're done
				m.quitting = true
				m.finalChoices = &Choices{
					IDEs:    []string{},
					Toolset: m.toolsetOptions[m.cursor],
				}
				for ide, selected := range m.selectedIDEs {
					if selected {
						m.finalChoices.IDEs = append(m.finalChoices.IDEs, ide)
					}
				}
				return m, tea.Quit
			}

		case " ":
			if m.step == stepIDE {
				// Toggle selection
				ide := m.ideOptions[m.cursor]
				m.selectedIDEs[ide] = !m.selectedIDEs[ide]
			}
		}
	}

	return m, nil
}

// View renders the UI.
func (m model) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder

	if m.step == stepIDE {
		b.WriteString("Which IDEs would you like to include?\n")
		b.WriteString("(Use space to select, enter to confirm)\n\n")
		for i, choice := range m.ideOptions {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			checked := " "
			if m.selectedIDEs[choice] {
				checked = "x"
			}
			b.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice))
		}
	} else {
		b.WriteString("Which toolset would you like to use?\n\n")
		for i, choice := range m.toolsetOptions {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}
	}

	b.WriteString("\n(q to quit)\n")
	return b.String()
}

// Run starts the Bubble Tea program and returns the final choices.
func Run() (*Choices, error) {
	p := tea.NewProgram(initialModel())

	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}

	// The final model is returned, so we can access the choices.
	if model, ok := m.(model); ok {
		return model.finalChoices, nil
	}

	return nil, fmt.Errorf("could not get final choices from TUI")
}
