package multiSelect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lf-silva/fastTrack/internal/ui/program"
)

var (
	headerStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("#E9681A")).Bold(true)
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

type Model struct {
	question  string
	cursor    int
	choices   []string
	submitted bool
	exit      *bool
}

func InitialModel(header string, choices []string, program *program.Project) Model {
	return Model{
		choices:  choices,
		question: header,
		exit:     &program.Exit,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.submitted = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := fmt.Sprintf("%s\n", headerStyle.Render(m.question))
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
			choice = selectedItemStyle.Render(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nUse ↑/↓ to navigate, Enter to select, and 'q' to quit."
	return s
}

func (m Model) GetCursor() int {
	return m.cursor
}
