package multiSelect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lf-silva/fastTrack/internal/ui/program"
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
	s := fmt.Sprintf("%s\n\n", m.question)
	for i, choice := range m.choices {
		cursor := " " // No cursor by default
		if m.cursor == i {
			cursor = ">" // Cursor points to the current choice
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nUse ↑/↓ to navigate, Enter to select, and 'q' to quit."
	return s
}

func (m Model) GetCursor() int {
	return m.cursor
}
