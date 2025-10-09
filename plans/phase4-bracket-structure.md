# Phase 4: Bracket Structure & Logic

**Status:** ⏳ PLANNED
**Branch:** `feature/phase4-bracket-structure` (to be created)
**Estimated Completion:** TBD

## Overview
Build the core data structures and algorithms for tournament bracket management. This phase establishes the foundation for bracket visualization and match tracking.

## Goals
- Define clean data structures for brackets, matches, and players
- Implement bracket generation algorithm
- Handle seeding and bye assignments
- Create match tree linking logic
- Support non-power-of-2 participant counts

## Tasks

### 1. Data Structure Definitions ⏳
**File:** `tournament/bracket.go` (new)

```go
// Player represents a tournament participant
type Player struct {
    ID   int
    Name string
    Seed int
}

// Match represents a single matchup in the tournament
type Match struct {
    ID          int
    Round       int
    Position    int       // Position in round (0-indexed)
    Player1     *Player   // nil if bye or TBD
    Player2     *Player   // nil if bye or TBD
    Winner      *Player   // nil until match is complete
    NextMatchID int       // ID of match winner advances to
    IsBye       bool      // True if one player gets automatic advancement
}

// Bracket represents the entire tournament structure
type Bracket struct {
    Participants    []Player
    Matches         []Match
    TotalRounds     int
    BracketSize     int
    CurrentRound    int
    IsComplete      bool
}
```

### 2. Bracket Builder Core ⏳
**File:** `tournament/bracket_builder.go` (new)

Main builder function:
```go
// NewBracket creates a complete bracket structure from participant count
func NewBracket(participantCount int) *Bracket {
    // 1. Calculate bracket parameters
    // 2. Create players with default names
    // 3. Generate all matches
    // 4. Assign seeding
    // 5. Distribute byes
    // 6. Link matches (winner advancement)
    return bracket
}
```

### 3. Match Generation Algorithm ⏳

**Algorithm Steps:**
1. Calculate total matches needed:
   ```
   totalMatches = bracketSize - 1
   ```
   (e.g., 16-player bracket needs 15 matches: 8+4+2+1)

2. Generate matches round by round:
   ```go
   func generateMatches(bracketSize, totalRounds int) []Match {
       matches := []Match{}
       matchID := 0

       for round := 0; round < totalRounds; round++ {
           matchesInRound := bracketSize / int(math.Pow(2, float64(round+1)))

           for pos := 0; pos < matchesInRound; pos++ {
               match := Match{
                   ID:       matchID,
                   Round:    round,
                   Position: pos,
               }
               matches = append(matches, match)
               matchID++
           }
       }

       return matches
   }
   ```

3. Link matches (winner advancement):
   ```go
   func linkMatches(matches []Match, totalRounds int) {
       for _, match := range matches {
           if match.Round < totalRounds-1 {
               // Winner advances to:
               // round + 1, position = currentPosition / 2
               nextRound := match.Round + 1
               nextPosition := match.Position / 2
               match.NextMatchID = findMatchID(matches, nextRound, nextPosition)
           }
       }
   }
   ```

### 4. Seeding System ⏳

**Standard Bracket Seeding:**
For a 16-player bracket:
```
Match 1: Seed 1  vs Seed 16
Match 2: Seed 8  vs Seed 9
Match 3: Seed 4  vs Seed 13
Match 4: Seed 5  vs Seed 12
Match 5: Seed 2  vs Seed 15
Match 6: Seed 7  vs Seed 10
Match 7: Seed 3  vs Seed 14
Match 8: Seed 6  vs Seed 11
```

**Seeding Algorithm:**
```go
func generateSeedOrder(size int) []int {
    if size == 1 {
        return []int{1}
    }

    prev := generateSeedOrder(size / 2)
    seeds := []int{}

    for _, seed := range prev {
        seeds = append(seeds, seed)
        seeds = append(seeds, size+1-seed)
    }

    return seeds
}
```

### 5. Bye Distribution ⏳

**Strategy:** Top seeds get byes

```go
func assignByes(bracket *Bracket, byeCount int) {
    // Top `byeCount` seeds get byes
    for i := 0; i < byeCount; i++ {
        player := &bracket.Participants[i]

        // Find player's first round match
        match := findPlayerFirstMatch(bracket, player)

        // Mark as bye, assign winner
        match.IsBye = true
        match.Player1 = player
        match.Player2 = nil
        match.Winner = player

        // Advance player to next match
        advancePlayer(bracket, match.NextMatchID, player)
    }
}
```

### 6. Player Assignment ⏳

```go
func assignPlayers(bracket *Bracket) {
    seedOrder := generateSeedOrder(bracket.BracketSize)

    // Only assign real participants (skip bye positions)
    realParticipants := len(bracket.Participants)

    matchIdx := 0
    for i := 0; i < len(seedOrder); i += 2 {
        match := &bracket.Matches[matchIdx]

        if seedOrder[i] <= realParticipants {
            match.Player1 = &bracket.Participants[seedOrder[i]-1]
        }

        if seedOrder[i+1] <= realParticipants {
            match.Player2 = &bracket.Participants[seedOrder[i+1]-1]
        }

        matchIdx++
    }
}
```

## Deliverables

### New Files
```
tournament/
├── bracket.go              # Data structures
├── bracket_builder.go      # Bracket generation
└── bracket_test.go         # Unit tests
```

### Features
- ✅ Player/Match/Bracket data structures
- ✅ Bracket generation for any participant count (2-64)
- ✅ Proper seeding algorithm
- ✅ Bye distribution logic
- ✅ Match linking (winner advancement)
- ✅ Helper methods for bracket queries

## Helper Methods

```go
// Bracket query methods
func (b *Bracket) GetMatchesInRound(round int) []Match
func (b *Bracket) GetMatch(matchID int) *Match
func (b *Bracket) GetPlayer(playerID int) *Player
func (b *Bracket) GetNextMatch(currentMatchID int) *Match
func (b *Bracket) IsMatchComplete(matchID int) bool
func (b *Bracket) GetPendingMatches() []Match
func (b *Bracket) GetCompletedMatches() []Match
func (b *Bracket) AdvanceWinner(matchID int, winnerID int) error
```

## Testing Strategy

### Unit Tests
Test with various participant counts:
- Powers of 2: 2, 4, 8, 16, 32, 64
- Non-powers: 3, 5, 6, 7, 12, 13, 20, 24, 48

### Test Cases
```go
func TestBracketGeneration(t *testing.T) {
    testCases := []struct{
        participants int
        expectedRounds int
        expectedByes int
        expectedMatches int
    }{
        {2, 1, 0, 1},
        {3, 2, 1, 3},
        {4, 2, 0, 3},
        {8, 3, 0, 7},
        {12, 4, 4, 15},
        {16, 4, 0, 15},
    }

    for _, tc := range testCases {
        bracket := NewBracket(tc.participants)
        assert.Equal(t, tc.expectedRounds, bracket.TotalRounds)
        assert.Equal(t, tc.expectedByes, countByes(bracket))
        assert.Equal(t, tc.expectedMatches, len(bracket.Matches))
    }
}
```

### Integration Tests
- Verify all matches are properly linked
- Ensure seeding is correct
- Validate bye assignments
- Test winner advancement logic

## Example Bracket Output

**8 Players, No Byes:**
```
Round 1 (4 matches):
  Match 0: Player 1 vs Player 8 → Winner to Match 4
  Match 1: Player 4 vs Player 5 → Winner to Match 4
  Match 2: Player 2 vs Player 7 → Winner to Match 5
  Match 3: Player 3 vs Player 6 → Winner to Match 5

Round 2 (2 matches):
  Match 4: TBD vs TBD → Winner to Match 6
  Match 5: TBD vs TBD → Winner to Match 6

Round 3 (1 match):
  Match 6: TBD vs TBD (Final)
```

**6 Players, 2 Byes:**
```
Round 1 (2 matches):
  Match 0: Player 3 vs Player 6 → Winner to Match 2
  Match 1: Player 4 vs Player 5 → Winner to Match 3

Round 2 (2 matches):
  Match 2: Player 1 (BYE) vs TBD → Winner to Match 4
  Match 3: Player 2 (BYE) vs TBD → Winner to Match 4

Round 3 (1 match):
  Match 4: TBD vs TBD (Final)
```

## Edge Cases

### 2 Participants (Minimum)
- 1 round
- 1 match
- No byes
- Simplest possible bracket

### 3 Participants
- 2 rounds
- 3 matches total
- 1 bye (top seed)
- Seed 1 gets bye to final
- Seed 2 vs Seed 3 in R1, winner faces Seed 1

### 64 Participants (Maximum)
- 6 rounds
- 63 matches total
- No byes (power of 2)
- Perfect bracket

## Next Steps
→ **Phase 5:** Bracket Visualization
- ASCII art rendering of bracket
- Multi-round display
- Color-coded match states
- Interactive navigation
