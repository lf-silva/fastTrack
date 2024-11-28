package multiSelect

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	question  string
	cursor    int
	choices   []string
	selected  map[int]struct{}
	submitted bool
}

func InitialModel(title string, choices []string) Model {
	return Model{
		choices:  choices,
		question: title,
		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.SetWindowTitle(m.question)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
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
			_, ok := m.selected[m.cursor]
			m.submitted = true
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	if m.submitted {
		// m.Answer = m.cursor
		return fmt.Sprintf("You chose: %d\nPress 'q' to quit.\n", m.cursor) // m.choices[m.cursor]
	}

	// Render the question and choices
	s := "What is your answer?\n\n"
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
