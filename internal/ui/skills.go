package ui

import tea "github.com/charmbracelet/bubbletea"

type skillsModel struct{}

func NewSkillsModel() tea.Model {
	return skillsModel{}
}

func (m skillsModel) Init() tea.Cmd {
	return nil
}

func (m skillsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m skillsModel) View() string {
	return "Skills page (press q to quit)\n"
}
