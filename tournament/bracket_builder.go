package tournament

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
