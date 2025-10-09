package tournament

// Player represents a tournament participant with seeding information.
type Player struct {
	ID   int    // Unique identifier for the player
	Name string // Display name of the player
	Seed int    // Seeding position (1 is highest seed)
}

// Match represents a single matchup in the tournament bracket.
// A match can be in various states: unplayed (both players TBD), bye (one player auto-advances),
// or completed (winner is set).
type Match struct {
	ID          int     // Unique identifier for the match
	Round       int     // Round number (0-indexed, 0 is first round)
	Position    int     // Position within the round (0-indexed)
	Player1     *Player // First player (nil if TBD or bye)
	Player2     *Player // Second player (nil if TBD or bye)
	Winner      *Player // Winner of the match (nil until match is complete)
	NextMatchID int     // ID of match winner advances to (-1 if final)
	IsBye       bool    // True if one player gets automatic advancement
}

// Bracket represents the complete tournament structure including all participants,
// matches, and tournament state.
type Bracket struct {
	Participants []Player // All tournament participants with seeding
	Matches      []Match  // All matches in the tournament
	TotalRounds  int      // Number of rounds in the tournament
	BracketSize  int      // Bracket size (next power of 2 from participant count)
	CurrentRound int      // Current active round (0-indexed)
	IsComplete   bool     // True when tournament has a winner
}
