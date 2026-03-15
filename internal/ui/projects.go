package ui

import tea "github.com/charmbracelet/bubbletea"

type projectsModel struct{}

func NewProjectsModel() tea.Model {
	return projectsModel{}
}

func (m projectsModel) Init() tea.Cmd {
	return nil
}

func (m projectsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m projectsModel) View() string {
	return "Projects page (press q to quit)\n"
}
