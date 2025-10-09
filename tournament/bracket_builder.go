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
