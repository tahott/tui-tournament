package tournament

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SEState int

const (
	SEStateSetup SEState = iota
	SEStateBracketView
	SEStateMatchEntry
)

type SingleEliminationModel struct {
	state            SEState
	participantCount int
	minParticipants  int
	maxParticipants  int
	width            int
	height           int
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

	seCountStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#4ECDC4")).
			Align(lipgloss.Center)

	seInfoBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#626262")).
			Padding(1, 2).
			Align(lipgloss.Center)

	seWarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD93D")).
			Italic(true).
			Align(lipgloss.Center)

	seLimitStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Align(lipgloss.Center)
)

func NewSingleEliminationModel() SingleEliminationModel {
	return SingleEliminationModel{
		state:            SEStateSetup,
		participantCount: 8,
		minParticipants:  2,
		maxParticipants:  64,
	}
}

func (m SingleEliminationModel) Init() tea.Cmd {
	return nil
}

func (m SingleEliminationModel) Update(msg tea.Msg) (SingleEliminationModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch m.state {
		case SEStateSetup:
			switch msg.String() {
			case "+", "j", "up":
				if m.participantCount < m.maxParticipants {
					m.participantCount++
				}
			case "-", "k", "down":
				if m.participantCount > m.minParticipants {
					m.participantCount--
				}
			case "enter":
				// Validate and transition to bracket view
				if m.participantCount >= m.minParticipants && m.participantCount <= m.maxParticipants {
					m.state = SEStateBracketView
				}
			}

		case SEStateBracketView:
			switch msg.String() {
			case "esc":
				// Return to setup
				m.state = SEStateSetup
			}
		}
	}
	return m, nil
}

func (m SingleEliminationModel) View() string {
	switch m.state {
	case SEStateSetup:
		return m.renderSetupView()
	case SEStateBracketView:
		return m.renderBracketView()
	case SEStateMatchEntry:
		// TODO: Phase 4 - Match entry view
		return m.renderBracketView()
	default:
		return m.renderSetupView()
	}
}

func (m SingleEliminationModel) renderSetupView() string {
	header := seHeaderStyle.Render("🥊 Single Elimination Tournament")

	// Calculate bracket properties
	rounds := CalculateRounds(m.participantCount)
	bracketSize := CalculateBracketSize(m.participantCount)
	byes := CalculateByes(m.participantCount)
	matchesPerRound := CalculateMatchesPerRound(m.participantCount)

	// Participant count display (large and prominent)
	countText := fmt.Sprintf("Participants: %d", m.participantCount)
	countDisplay := seCountStyle.Render(countText)

	// Show limits feedback
	var limitMsg string
	if m.participantCount == m.minParticipants {
		limitMsg = seLimitStyle.Render(fmt.Sprintf("(minimum: %d)", m.minParticipants))
	} else if m.participantCount == m.maxParticipants {
		limitMsg = seLimitStyle.Render(fmt.Sprintf("(maximum: %d)", m.maxParticipants))
	}

	// Bracket info section
	var infoLines []string
	infoLines = append(infoLines, fmt.Sprintf("Bracket Size: %d", bracketSize))
	infoLines = append(infoLines, fmt.Sprintf("Rounds: %d", rounds))
	infoLines = append(infoLines, fmt.Sprintf("Byes: %d", byes))
	infoLines = append(infoLines, "")

	// Round breakdown
	for i, matches := range matchesPerRound {
		roundNum := i + 1
		roundName := GetRoundName(roundNum, rounds)

		var line string
		if roundName != "" {
			line = fmt.Sprintf("Round %d: %d matches (%s)", roundNum, matches, roundName)
		} else {
			line = fmt.Sprintf("Round %d: %d matches", roundNum, matches)
		}

		// Add player count for round 1
		if i == 0 && byes > 0 {
			playersInRound1 := m.participantCount - byes
			line = fmt.Sprintf("Round %d: %d matches (%d players)", roundNum, matches, playersInRound1)
		}

		infoLines = append(infoLines, line)
	}

	infoBox := seInfoBoxStyle.Render(lipgloss.JoinVertical(lipgloss.Left, infoLines...))

	// Bye warning
	var byeWarning string
	if byes > 0 {
		byeWarning = seWarningStyle.Render(fmt.Sprintf("[%d players get byes]", byes))
	}

	// Help text
	help := seHelpStyle.Render("+ - or j k to adjust • Enter to continue • Esc to go back")

	// Combine all sections
	sections := []string{header, "", countDisplay}

	if limitMsg != "" {
		sections = append(sections, limitMsg)
	}

	sections = append(sections, "", infoBox)

	if byeWarning != "" {
		sections = append(sections, "", byeWarning)
	}

	sections = append(sections, "", help)

	view := lipgloss.JoinVertical(lipgloss.Center, sections...)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		view,
	)
}

func (m SingleEliminationModel) renderBracketView() string {
	header := seHeaderStyle.Render("🥊 Single Elimination Tournament")

	// Display current configuration
	configText := fmt.Sprintf("Tournament with %d participants", m.participantCount)
	rounds := CalculateRounds(m.participantCount)
	roundsText := fmt.Sprintf("%d rounds", rounds)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		"",
		configText,
		roundsText,
		"",
		"Bracket view coming in Phase 4...",
		"",
	)

	help := seHelpStyle.Render("Press Esc to go back to setup")

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
