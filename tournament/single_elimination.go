package tournament

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SingleEliminationModel struct {
	width  int
	height int
}

var (
	seHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6B6B")).
			Align(lipgloss.Center).
			MarginBottom(2)

	seHelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Align(lipgloss.Center).
			MarginTop(2)
)

func NewSingleEliminationModel() SingleEliminationModel {
	return SingleEliminationModel{}
}

func (m SingleEliminationModel) Init() tea.Cmd {
	return nil
}

func (m SingleEliminationModel) Update(msg tea.Msg) (SingleEliminationModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m SingleEliminationModel) View() string {
	header := seHeaderStyle.Render("🥊 Single Elimination Tournament")
	content := fmt.Sprintf("\n\nSingle Elimination Mode\n\nComing soon...\n\n")
	help := seHelpStyle.Render("Press Esc to go back to menu")

	view := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		content,
		help,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		view,
	)
}
