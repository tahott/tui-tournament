# Phase 5: Bracket Visualization

**Status:** ⏳ PLANNED
**Branch:** `feature/phase5-bracket-visualization` (to be created)
**Estimated Completion:** TBD

## Overview
Create a beautiful, functional ASCII art visualization of the tournament bracket that updates in real-time as matches are completed.

## Goals
- Display full tournament bracket in terminal
- Show all rounds side-by-side
- Color-code match states (pending, in-progress, completed)
- Handle various bracket sizes gracefully
- Support navigation and scrolling for large brackets
- Highlight current/selected match

## Tasks

### 1. Bracket Renderer Core ⏳
**File:** `tournament/bracket_renderer.go` (new)

```go
type BracketRenderer struct {
    bracket      *Bracket
    width        int
    height       int
    scrollOffset int
    selectedMatch int
}

func (r *BracketRenderer) Render() string {
    // Main rendering function
    // Returns full bracket as styled string
}
```

### 2. ASCII Art Layout ⏳

**Layout Strategy:**
- Vertical display: rounds go left-to-right
- Each match is a box with player names
- Lines connect matches to show progression
- Spacing increases between matches in later rounds

**Example 8-Player Bracket:**
```
Round 1          Round 2        Final

┌─────────────┐
│ Player 1    │──┐
│ vs          │  │  ┌─────────────┐
│ Player 8    │──┘  │ Winner 1    │──┐
└─────────────┘     │ vs          │  │
                    │ Winner 2    │──┘  ┌─────────────┐
┌─────────────┐     └─────────────┘     │ Champion    │
│ Player 4    │──┐                       │             │
│ vs          │  │                       └─────────────┘
│ Player 5    │──┘
└─────────────┘

┌─────────────┐
│ Player 2    │──┐
│ vs          │  │  ┌─────────────┐
│ Player 7    │──┘  │ Winner 3    │──┐
└─────────────┘     │ vs          │  │
                    │ Winner 4    │──┘
┌─────────────┐     └─────────────┘
│ Player 3    │──┐
│ vs          │  │
│ Player 6    │──┘
└─────────────┘
```

### 3. Match Box Component ⏳

```go
func renderMatchBox(match Match, state MatchState, isSelected bool) string {
    var style lipgloss.Style

    switch state {
    case MatchStatePending:
        style = pendingMatchStyle
    case MatchStateInProgress:
        style = activeMatchStyle
    case MatchStateCompleted:
        style = completedMatchStyle
    }

    if isSelected {
        style = style.Copy().Border(lipgloss.DoubleBorder())
    }

    player1 := "TBD"
    player2 := "TBD"

    if match.Player1 != nil {
        player1 = match.Player1.Name
    }
    if match.Player2 != nil {
        player2 = match.Player2.Name
    }

    content := fmt.Sprintf("%s\nvs\n%s", player1, player2)

    if match.Winner != nil {
        // Highlight winner
        if match.Winner == match.Player1 {
            content = fmt.Sprintf("✓ %s\n  vs\n  %s", player1, player2)
        } else {
            content = fmt.Sprintf("  %s\n  vs\n✓ %s", player1, player2)
        }
    }

    return style.Render(content)
}
```

### 4. Connection Lines ⏳

**Rendering connector lines between matches:**

```go
func renderConnector(fromY, toY int, length int) string {
    // Horizontal line
    if fromY == toY {
        return strings.Repeat("─", length)
    }

    // Vertical line with junction
    lines := []string{}
    for y := fromY; y <= toY; y++ {
        if y == fromY {
            lines = append(lines, "──┐")
        } else if y == toY {
            lines = append(lines, "  │")
        } else {
            lines = append(lines, "  │")
        }
    }

    return strings.Join(lines, "\n")
}
```

### 5. Round Column Layout ⏳

```go
func (r *BracketRenderer) renderRound(round int) string {
    matches := r.bracket.GetMatchesInRound(round)
    roundColumn := []string{}

    // Add round header
    roundColumn = append(roundColumn, renderRoundHeader(round))

    // Spacing increases by round (2^round)
    spacing := int(math.Pow(2, float64(round)))

    for i, match := range matches {
        // Add spacing
        for j := 0; j < spacing-1; j++ {
            roundColumn = append(roundColumn, "")
        }

        // Add match box
        matchBox := renderMatchBox(match, r.getMatchState(match), r.selectedMatch == match.ID)
        roundColumn = append(roundColumn, matchBox)

        // Add spacing
        for j := 0; j < spacing; j++ {
            roundColumn = append(roundColumn, "")
        }
    }

    return lipgloss.JoinVertical(lipgloss.Left, roundColumn...)
}
```

### 6. Full Bracket Assembly ⏳

```go
func (r *BracketRenderer) Render() string {
    rounds := []string{}

    for round := 0; round < r.bracket.TotalRounds; round++ {
        roundView := r.renderRound(round)
        rounds = append(rounds, roundView)
    }

    // Join rounds horizontally with connectors
    bracket := lipgloss.JoinHorizontal(lipgloss.Top, rounds...)

    return bracket
}
```

### 7. Color Coding ⏳

**Match States:**
```go
var (
    pendingMatchStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#666666")).
        Foreground(lipgloss.Color("#AAAAAA")).
        Padding(0, 1)

    activeMatchStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#FFD700")).
        Foreground(lipgloss.Color("#FFFFFF")).
        Background(lipgloss.Color("#3A3A00")).
        Padding(0, 1)

    completedMatchStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#00FF00")).
        Foreground(lipgloss.Color("#CCCCCC")).
        Padding(0, 1)

    selectedMatchStyle = lipgloss.NewStyle().
        Border(lipgloss.DoubleBorder()).
        BorderForeground(lipgloss.Color("#FF69B4")).
        Padding(0, 1)
)
```

### 8. Navigation & Scrolling ⏳

**For Large Brackets:**
```go
func (r *BracketRenderer) Update(msg tea.Msg) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "up", "k":
            r.selectPreviousMatch()
        case "down", "j":
            r.selectNextMatch()
        case "left", "h":
            r.selectMatchInPreviousRound()
        case "right", "l":
            r.selectMatchInNextRound()
        case "pgup":
            r.scrollUp()
        case "pgdown":
            r.scrollDown()
        }
    }
}
```

### 9. Responsive Design ⏳

Handle different terminal sizes:
- **Large terminals (>120 cols):** Full horizontal layout
- **Medium terminals (80-120 cols):** Compressed layout
- **Small terminals (<80 cols):** Vertical scrolling, show 2 rounds at a time

```go
func (r *BracketRenderer) getLayoutMode() LayoutMode {
    if r.width > 120 {
        return LayoutModeFull
    } else if r.width > 80 {
        return LayoutModeCompact
    } else {
        return LayoutModeScrollable
    }
}
```

### 10. Round Headers ⏳

```go
func renderRoundHeader(round, totalRounds int) string {
    var label string

    if round == totalRounds-1 {
        label = "FINAL"
    } else if round == totalRounds-2 {
        label = "SEMI-FINAL"
    } else {
        label = fmt.Sprintf("Round %d", round+1)
    }

    return headerStyle.Render(label)
}
```

## Deliverables

### New Files
```
tournament/
├── bracket_renderer.go     # Bracket visualization logic
└── styles.go               # Lipgloss styles for bracket
```

### Features
- ✅ Full bracket ASCII art rendering
- ✅ Side-by-side round display
- ✅ Color-coded match states
- ✅ Winner highlighting (✓ symbol)
- ✅ Connection lines between matches
- ✅ Round headers with labels
- ✅ Selected match highlighting
- ✅ Keyboard navigation
- ✅ Scrolling for large brackets
- ✅ Responsive to terminal size

## Visual Examples

### 4-Player Bracket
```
   Round 1              Final

┌──────────────┐
│ ✓ Player 1   │──┐
│   vs         │  │
│   Player 4   │  │    ┌──────────────┐
└──────────────┘  └────│   Player 1   │
                       │      vs      │
┌──────────────┐  ┌────│     TBD      │
│   Player 2   │  │    └──────────────┘
│   vs         │  │
│ ✓ Player 3   │──┘
└──────────────┘
```

### With Byes (6 Players)
```
   Round 1         Round 2           Final

┌──────────────┐
│   Player 3   │──┐
│   vs         │  │  ┌──────────────┐
│ ✓ Player 6   │──┘  │ Player 1(BYE)│──┐
└──────────────┘     │      vs      │  │
                     │   Player 6   │  │  ┌──────────────┐
                     └──────────────┘  └──│     TBD      │
                                          │      vs      │
┌──────────────┐     ┌──────────────┐  ┌──│     TBD      │
│   Player 4   │──┐  │ Player 2(BYE)│  │  └──────────────┘
│   vs         │  │  │      vs      │──┘
│   Player 5   │──┘  │     TBD      │
└──────────────┘     └──────────────┘
```

## Interaction Flow

1. **View Mode (Default):**
   - Display full bracket
   - No match selected
   - Show all match states

2. **Selection Mode:**
   - Press j/k or ↑/↓ to select match
   - Selected match has double border
   - Can navigate between rounds with h/l or ←/→

3. **Enter Match Result:**
   - Press Enter on selected match
   - Opens result entry (Phase 6)

## Testing

### Visual Tests
Test rendering with:
- 2, 4, 8, 16, 32 participants (powers of 2)
- 3, 5, 6, 7, 12, 24 participants (with byes)
- All matches pending
- Some matches completed
- All matches completed (champion)

### Edge Cases
- Very long player names (truncation)
- Very small terminal size
- Very large bracket (64 players)
- Single match (2 players)

## Performance Considerations
- Cache rendered brackets when state doesn't change
- Use viewport for large brackets to limit rendering
- Optimize string concatenation
- Lazy render: only render visible portion in scroll mode

## Next Steps
→ **Phase 6:** Match Entry & Progression
- Select match to update
- Choose winner
- Automatic bracket updates
- Tournament completion detection
