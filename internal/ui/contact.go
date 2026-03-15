package ui

import tea "github.com/charmbracelet/bubbletea"

type contactModel struct{}

func NewContactModel() tea.Model {
	return contactModel{}
}

func (m contactModel) Init() tea.Cmd {
	return nil
}

func (m contactModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m contactModel) View() string {
	return "Contact page (press q to quit)\n"
}
