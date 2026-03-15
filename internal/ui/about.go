package ui

import tea "github.com/charmbracelet/bubbletea"

type aboutModel struct{}

func NewAboutModel() tea.Model {
	return aboutModel{}
}

func (m aboutModel) Init() tea.Cmd {
	return nil
}

func (m aboutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m aboutModel) View() string {
	return "About page (press q to quit)\n"
}
