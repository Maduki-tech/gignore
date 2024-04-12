package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	programmingLanguages := []string{
		"Java",
		"Go",
		"Python",
		"JavaScript",
		"TypeScript",
	}
	return model{
		choices:  programmingLanguages,
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

		case "enter":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				// m.selected[m.cursor] = struct{}{}
				DownloadIgnore(m.choices[m.cursor])
				return m, tea.Quit
			}

		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Which .gitignore file would you like to download?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += cursor + " " + checked + " " + choice + "\n"
	}

	s += "\nPress q to quit. Press enter to select/unselect a file.\n\n"

	return s
}
