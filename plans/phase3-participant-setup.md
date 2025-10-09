# Phase 3: Single Elimination - Participant Setup

**Status:** 🚧 IN PROGRESS
**Branch:** `feature/phase3-participant-setup` (to be created)
**Estimated Completion:** TBD

## Overview
Implement the participant configuration screen for Single Elimination tournaments. Users can adjust the number of participants and see a real-time preview of how the bracket will be structured.

## Goals
- Allow users to configure tournament size (2-64 participants)
- Display dynamic bracket preview that updates in real-time
- Show bracket structure information (rounds, byes, etc.)
- Validate participant count and provide clear feedback
- Enable smooth transition to bracket generation

## Tasks

### 1. Participant Count Model ⏳
**File:** `tournament/single_elimination.go`

Add state to track participant configuration:
```go
type SingleEliminationModel struct {
    state            SEState  // Setup, BracketView, MatchEntry
    participantCount int
    minParticipants  int      // 2
    maxParticipants  int      // 64
    width            int
    height           int
}

type SEState int
const (
    SEStateSetup SEState = iota
    SEStateBracketView
    SEStateMatchEntry
)
```

### 2. Count Adjustment Controls ⏳
Implement keyboard controls for adjusting participant count:
- **+** or **j** or **↑**: Increase participant count
- **-** or **k** or **↓**: Decrease participant count
- **Enter**: Proceed to bracket generation
- **Esc**: Return to menu

**Validation:**
- Don't allow count < 2
- Don't allow count > 64
- Provide visual feedback when at limits

### 3. Bracket Calculation Logic ⏳
**File:** `tournament/bracket_calculator.go` (new)

Create utility functions to calculate bracket properties:
```go
// Calculate number of rounds needed
func CalculateRounds(participants int) int {
    return int(math.Ceil(math.Log2(float64(participants))))
}

// Calculate bracket size (next power of 2)
func CalculateBracketSize(participants int) int {
    return int(math.Pow(2, math.Ceil(math.Log2(float64(participants)))))
}

// Calculate number of byes needed
func CalculateByes(participants int) int {
    return CalculateBracketSize(participants) - participants
}

// Calculate matches per round
func CalculateMatchesPerRound(participants int) []int {
    // Returns slice of match counts for each round
}
```

### 4. Bracket Preview Visualization ⏳
Create a visual preview showing:
- Current participant count (large, prominent)
- Number of rounds
- Bracket size (next power of 2)
- Number of byes (if applicable)
- Simple ASCII bracket structure preview

**Example Preview:**
```
┌─────────────────────────────────────┐
│     Tournament Configuration        │
├─────────────────────────────────────┤
│                                     │
│        Participants: 12             │
│                                     │
│        Bracket Size: 16             │
│        Rounds: 4                    │
│        Byes: 4                      │
│                                     │
│   Round 1: 4 matches (8 players)   │
│   Round 2: 4 matches                │
│   Round 3: 2 matches                │
│   Round 4: 1 match (Final)          │
│                                     │
│        [4 players get byes]         │
│                                     │
├─────────────────────────────────────┤
│  + - or j k to adjust • Enter to    │
│  continue • Esc to go back          │
└─────────────────────────────────────┘
```

### 5. Dynamic Preview Updates ⏳
- Update all calculations immediately when count changes
- Smooth visual feedback
- Highlight when at min/max limits
- Show warnings for edge cases (e.g., "All players get byes except 2")

### 6. Visual Design ⏳
**Styling with Lipgloss:**
- Large, bold participant count display
- Info box for bracket details
- Color-coded elements:
  - Participant count: Bright/prominent color
  - Bracket info: Neutral color
  - Byes warning: Yellow/orange (if > 0)
  - Limits reached: Red
- Centered layout
- Responsive to terminal size

### 7. Proceed to Next Stage ⏳
When user presses Enter:
- Validate participant count
- Transition to `SEStateBracketView` state
- Pass participant count to bracket generation
- Initialize bracket structure (Phase 4 work)

## Deliverables

### New Files
```
tournament/
├── single_elimination.go      # Updated with setup state
└── bracket_calculator.go      # New - bracket math utilities
```

### Features
- ✅ Participant count adjustment (2-64)
- ✅ Real-time bracket preview
- ✅ Round and bye calculations
- ✅ Visual feedback for limits
- ✅ Smooth keyboard controls
- ✅ Transition to bracket generation

## Technical Details

### State Machine
```
SEStateSetup → (Enter) → SEStateBracketView → (Enter) → SEStateMatchEntry
     ↑
     └────────────────── (Esc) ──────────────────────────────┘
```

### Bracket Size Examples
| Participants | Bracket Size | Rounds | Byes |
|--------------|--------------|--------|------|
| 2            | 2            | 1      | 0    |
| 3            | 4            | 2      | 1    |
| 4            | 4            | 2      | 0    |
| 5            | 8            | 3      | 3    |
| 8            | 8            | 3      | 0    |
| 12           | 16           | 4      | 4    |
| 16           | 16           | 4      | 0    |
| 32           | 32           | 5      | 0    |

### Bye Distribution Strategy
For initial implementation, byes go to top seeds:
- If 12 participants (4 byes needed)
- Seeds 1, 2, 3, 4 get byes to Round 2
- Seeds 5-12 play in Round 1

## User Experience Flow
1. User selects "Single Elimination" from menu
2. Screen shows participant setup with default count (8)
3. User adjusts count up/down
4. Preview updates in real-time
5. User satisfied with configuration
6. Presses Enter to proceed
7. → Transitions to bracket view (Phase 4)

## Edge Cases to Handle
- ✅ Minimum participants (2) - disable decrease
- ✅ Maximum participants (64) - disable increase
- ✅ Power of 2 counts (4, 8, 16, 32, 64) - no byes needed
- ✅ All byes except 2 players (e.g., 3 participants)
- ✅ Very small terminal sizes

## Testing Checklist
- [ ] Count increases correctly
- [ ] Count decreases correctly
- [ ] Can't go below 2
- [ ] Can't go above 64
- [ ] Calculations correct for all values 2-64
- [ ] Visual preview matches calculations
- [ ] Byes calculated correctly
- [ ] Rounds calculated correctly
- [ ] Enter proceeds to next state
- [ ] Esc returns to menu
- [ ] Responsive to window resizing

## Next Steps
→ **Phase 4:** Bracket Structure & Logic
- Define data structures (Match, Player, Bracket)
- Implement bracket generation algorithm
- Create bracket builder from participant count
