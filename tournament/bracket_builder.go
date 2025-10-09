package tournament

import "fmt"

// NewBracket creates a complete bracket structure from participant count.
// It orchestrates all bracket generation steps: creating players, generating matches,
// applying seeding, distributing byes, and linking matches for winner advancement.
func NewBracket(participantCount int) *Bracket {
	// 1. Calculate bracket parameters
	rounds := CalculateRounds(participantCount)
	bracketSize := CalculateBracketSize(participantCount)
	byes := CalculateByes(participantCount)

	// 2. Create players with default names and seeding
	participants := make([]Player, participantCount)
	for i := 0; i < participantCount; i++ {
		participants[i] = Player{
			ID:   i,
			Name: fmt.Sprintf("Player %d", i+1),
			Seed: i + 1, // Seed 1 is highest
		}
	}

	// 3. Generate all matches
	matches := generateMatches(bracketSize, rounds)

	// 4. Link matches (winner advancement)
	linkMatches(matches, rounds)

	// 5. Initialize bracket
	bracket := &Bracket{
		Participants: participants,
		Matches:      matches,
		TotalRounds:  rounds,
		BracketSize:  bracketSize,
		CurrentRound: 0,
		IsComplete:   false,
	}

	// 6. Assign seeding to matches
	assignPlayers(bracket)

	// 7. Distribute byes to top seeds
	if byes > 0 {
		assignByes(bracket, byes)
	}

	return bracket
}

// assignPlayers assigns players to first round matches based on standard bracket seeding.
func assignPlayers(bracket *Bracket) {
	seedOrder := generateSeedOrder(bracket.BracketSize)
	realParticipants := len(bracket.Participants)

	matchIdx := 0
	for i := 0; i < len(seedOrder); i += 2 {
		if matchIdx >= len(bracket.Matches) {
			break
		}

		match := &bracket.Matches[matchIdx]

		// Only assign if seed exists (handles non-power-of-2)
		if seedOrder[i] <= realParticipants {
			match.Player1 = &bracket.Participants[seedOrder[i]-1]
		}

		if seedOrder[i+1] <= realParticipants {
			match.Player2 = &bracket.Participants[seedOrder[i+1]-1]
		}

		matchIdx++
	}
}

// generateMatches creates all matches for a tournament bracket based on bracket size and total rounds.
// Matches are created round by round, with each round having half the matches of the previous round.
// Returns a slice of matches with IDs, round numbers, and positions assigned.
func generateMatches(bracketSize, totalRounds int) []Match {
	matches := []Match{}
	matchID := 0

	for round := 0; round < totalRounds; round++ {
		// Calculate matches in this round: bracketSize / 2^(round+1)
		matchesInRound := bracketSize >> (round + 1)

		for pos := 0; pos < matchesInRound; pos++ {
			match := Match{
				ID:          matchID,
				Round:       round,
				Position:    pos,
				NextMatchID: -1, // Will be set by linkMatches
			}
			matches = append(matches, match)
			matchID++
		}
	}

	return matches
}

// findMatchID finds the match ID for a given round and position.
// Returns -1 if not found.
func findMatchID(matches []Match, round, position int) int {
	for _, match := range matches {
		if match.Round == round && match.Position == position {
			return match.ID
		}
	}
	return -1
}

// linkMatches establishes winner advancement paths by setting NextMatchID for each match.
// Winners advance to the next round at position = currentPosition / 2.
// The final match has NextMatchID = -1 (no advancement).
func linkMatches(matches []Match, totalRounds int) {
	for i := range matches {
		match := &matches[i]

		// Final match has no next match
		if match.Round >= totalRounds-1 {
			match.NextMatchID = -1
			continue
		}

		// Winner advances to next round, position divided by 2
		nextRound := match.Round + 1
		nextPosition := match.Position / 2
		match.NextMatchID = findMatchID(matches, nextRound, nextPosition)
	}
}

// generateSeedOrder creates the standard bracket seeding order recursively.
// For a bracket size of n, returns a slice of seed numbers in the order they should appear
// in the bracket to ensure proper competitive balance (top seeds don't meet until later rounds).
//
// Example for size 16:
//   [1, 16, 8, 9, 4, 13, 5, 12, 2, 15, 7, 10, 3, 14, 6, 11]
//
// This creates matchups: 1v16, 8v9, 4v13, 5v12, 2v15, 7v10, 3v14, 6v11
func generateSeedOrder(size int) []int {
	if size == 1 {
		return []int{1}
	}

	// Get seeding for half bracket
	prev := generateSeedOrder(size / 2)
	seeds := []int{}

	// For each seed in previous level, add it and its complement
	for _, seed := range prev {
		seeds = append(seeds, seed)
		seeds = append(seeds, size+1-seed)
	}

	return seeds
}

// findPlayerFirstMatch finds the first match containing the given player.
// Returns nil if player is not found in any match.
func findPlayerFirstMatch(bracket *Bracket, player *Player) *Match {
	for i := range bracket.Matches {
		match := &bracket.Matches[i]
		if match.Player1 == player || match.Player2 == player {
			return match
		}
	}
	return nil
}

// advancePlayer places a player in the next match after winning.
// Assigns to the first empty player slot (Player1 or Player2).
func advancePlayer(bracket *Bracket, nextMatchID int, player *Player) {
	if nextMatchID == -1 {
		return
	}

	nextMatch := &bracket.Matches[nextMatchID]

	// Assign to first empty slot
	if nextMatch.Player1 == nil {
		nextMatch.Player1 = player
	} else if nextMatch.Player2 == nil {
		nextMatch.Player2 = player
	}
}

// assignByes assigns automatic advancement (byes) to top-seeded players.
// Top seeds skip round 1 when participant count is not a power of 2.
// The number of byes = bracketSize - participantCount.
func assignByes(bracket *Bracket, byeCount int) {
	// Top `byeCount` seeds get byes
	for i := 0; i < byeCount; i++ {
		player := &bracket.Participants[i]

		// Find player's first round match
		match := findPlayerFirstMatch(bracket, player)
		if match == nil {
			continue
		}

		// Mark as bye, assign winner
		match.IsBye = true
		match.Player1 = player
		match.Player2 = nil
		match.Winner = player

		// Advance player to next match
		advancePlayer(bracket, match.NextMatchID, player)
	}
}
