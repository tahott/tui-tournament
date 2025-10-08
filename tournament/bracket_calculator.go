package tournament

import "math/bits"

// CalculateRounds calculates the number of rounds needed for a given number of participants.
// For example:
//   - 2 participants -> 1 round
//   - 3-4 participants -> 2 rounds
//   - 5-8 participants -> 3 rounds
//   - 9-16 participants -> 4 rounds
func CalculateRounds(participants int) int {
	if participants <= 1 {
		return 0
	}
	// ceil(log2(n)) == bits.Len(n-1) for n>1
	return bits.Len(uint(participants - 1))
}

// CalculateBracketSize calculates the bracket size (next power of 2) for a given number of participants.
// For example:
//   - 2 participants -> 2
//   - 3 participants -> 4
//   - 5 participants -> 8
//   - 12 participants -> 16
func CalculateBracketSize(participants int) int {
	if participants <= 1 {
		return 1
	}
	rounds := CalculateRounds(participants)
	return 1 << rounds
}

// CalculateByes calculates the number of byes needed for a given number of participants.
// Byes are given to top seeds to balance the bracket.
// For example:
//   - 2 participants -> 0 byes
//   - 3 participants -> 1 bye
//   - 5 participants -> 3 byes
//   - 12 participants -> 4 byes
func CalculateByes(participants int) int {
	return CalculateBracketSize(participants) - participants
}

// CalculateMatchesPerRound calculates the number of matches in each round.
// Returns a slice where index 0 is Round 1, index 1 is Round 2, etc.
// For example, with 12 participants:
//   - Round 1: 4 matches (8 players, 4 get byes)
//   - Round 2: 4 matches (4 winners + 4 bye players)
//   - Round 3: 2 matches (semifinals)
//   - Round 4: 1 match (final)
func CalculateMatchesPerRound(participants int) []int {
	if participants <= 1 {
		return []int{}
	}

	rounds := CalculateRounds(participants)
	byes := CalculateByes(participants)

	matches := make([]int, rounds)

	// Round 1: Only non-bye players compete
	playersInRound1 := participants - byes
	matches[0] = playersInRound1 / 2

	// Subsequent rounds: half the players from previous round + any bye players joining
	remainingPlayers := matches[0] // winners from round 1
	if byes > 0 {
		remainingPlayers += byes // bye players join in round 2
	}

	for i := 1; i < rounds; i++ {
		matches[i] = remainingPlayers / 2
		remainingPlayers = matches[i]
	}

	return matches
}

// GetRoundName returns a human-readable name for a given round number.
// The last round is "Final", second-to-last is "Semifinals", etc.
func GetRoundName(roundNumber, totalRounds int) string {
	if roundNumber < 1 || roundNumber > totalRounds {
		return ""
	}

	roundsFromEnd := totalRounds - roundNumber

	switch roundsFromEnd {
	case 0:
		return "Final"
	case 1:
		return "Semifinals"
	case 2:
		return "Quarterfinals"
	default:
		return ""
	}
}
