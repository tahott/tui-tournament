package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"go-tournament/tournament"
)

type model struct {
	currentScreen       Screen
	menuModel           menuModel
	singleElimination   tournament.SingleEliminationModel
	width               int
	height              int
}

func newModel() model {
	return model{
		currentScreen:     ScreenMenu,
		menuModel:         newMenuModel(),
		singleElimination: tournament.NewSingleEliminationModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// switchScreen changes the current screen and injects window size to the new screen model
func (m *model) switchScreen(target Screen) {
	m.currentScreen = target
	// Pass current window size to new screen model to ensure proper initial rendering
	sizeMsg := tea.WindowSizeMsg{Width: m.width, Height: m.height}
	switch target {
	case ScreenMenu:
		m.menuModel, _ = m.menuModel.Update(sizeMsg)
	case ScreenSingleElimination:
		updated, _ := m.singleElimination.Update(sizeMsg)
		m.singleElimination = updated.(tournament.SingleEliminationModel)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			// Go back to menu from any screen
			if m.currentScreen != ScreenMenu {
				m.switchScreen(ScreenMenu)
				return m, nil
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case screenChangeMsg:
		// Handle screen changes
		m.switchScreen(msg.screen)
		return m, nil
	}

	// Delegate to the appropriate screen model
	var cmd tea.Cmd
	switch m.currentScreen {
	case ScreenMenu:
		m.menuModel, cmd = m.menuModel.Update(msg)
	case ScreenSingleElimination:
		updated, c := m.singleElimination.Update(msg)
		m.singleElimination = updated.(tournament.SingleEliminationModel)
		cmd = c
	}

	return m, cmd
}

func (m model) View() string {
	// Delegate to the appropriate screen view
	switch m.currentScreen {
	case ScreenMenu:
		return m.menuModel.View()
	case ScreenSingleElimination:
		return m.singleElimination.View()
	default:
		return "Unknown screen"
	}
}

func main() {
	m := newModel()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
