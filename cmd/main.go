package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/EbraamSobhy/Portfolio-CLI/internal/ui"
)

func main() {
    p := tea.NewProgram(ui.NewMenuModel())
    p.Start()
}
