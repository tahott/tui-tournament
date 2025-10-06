package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tournamentType struct {
	name        string
	description string
	icon        string
	screen      Screen
}

type menuModel struct {
	tournaments []tournamentType
	selected    int
	width       int
	height      int
}

var (
	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2).
			Width(25).
			Height(8)

	selectedCardStyle = cardStyle.Copy().
				BorderForeground(lipgloss.Color("#FF69B4")).
				Background(lipgloss.Color("#1A1A2E"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Align(lipgloss.Center).
			MarginBottom(1)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6B6B")).
			Align(lipgloss.Center).
			MarginBottom(2)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			Align(lipgloss.Center).
			MarginTop(2)
)

func newMenuModel() menuModel {
	return menuModel{
		tournaments: []tournamentType{
			{
				name:        "Single Elimination",
				description: "Classic bracket style\nwhere losers are\neliminated instantly",
				icon:        "🥊",
				screen:      ScreenSingleElimination,
			},
			{
				name:        "Double Elimination",
				description: "Players get a second\nchance in the\nloser's bracket",
				icon:        "🔄",
				screen:      ScreenDoubleElimination,
			},
			{
				name:        "Round Robin",
				description: "Everyone plays\neveryone else\nat least once",
				icon:        "🔁",
				screen:      ScreenRoundRobin,
			},
		},
		selected: 0,
	}
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

type screenChangeMsg struct {
	screen Screen
}

func (m menuModel) Update(msg tea.Msg) (menuModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			if m.selected > 0 {
				m.selected--
			}
		case "right", "l":
			if m.selected < len(m.tournaments)-1 {
				m.selected++
			}
		case "enter":
			// Send message to switch screen
			selectedScreen := m.tournaments[m.selected].screen
			return m, func() tea.Msg {
				return screenChangeMsg{screen: selectedScreen}
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m menuModel) View() string {
	var cards []string

	for i, tournament := range m.tournaments {
		var style lipgloss.Style
		if i == m.selected {
			style = selectedCardStyle
		} else {
			style = cardStyle
		}

		content := fmt.Sprintf("%s\n\n%s\n\n%s",
			tournament.icon,
			titleStyle.Render(tournament.name),
			tournament.description,
		)

		cards = append(cards, style.Render(content))
	}

	cardsRow := lipgloss.JoinHorizontal(lipgloss.Center, cards...)

	header := headerStyle.Render("🏆 Tournament Manager 🏆")
	help := helpStyle.Render("← → or h l to navigate • Enter to select • q to quit")

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		header,
		cardsRow,
		help,
	)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}
