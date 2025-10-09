# Go Tournament Manager - Development Plan Overview

## Project Vision
A beautiful, functional terminal-based tournament management application built with Go and Bubbletea, supporting multiple tournament formats.

## Development Phases

### Phase 1: Project Setup ✅ COMPLETED
**Status:** Completed
**Branch:** `main`
**Description:** Initial project setup with basic TUI framework

- Go module initialization
- Bubbletea & Lipgloss dependencies installation
- Tournament selection screen with card-based UI
- Navigation between tournament types
- README documentation

**Deliverables:**
- Working tournament selection menu
- 3 tournament type cards (Single Elimination, Double Elimination, Round Robin)
- Basic keyboard navigation

---

### Phase 2: Application Structure Refactoring ✅ COMPLETED
**Status:** Completed
**Branch:** `feature/phase2-app-structure-refactoring`
**Description:** Modular architecture for screen management

- Screen state management system
- Separation of menu and tournament screens
- Package structure for tournament logic
- Screen navigation (menu ↔ tournament screens)

**Deliverables:**
- `screen.go` - Screen state enum
- `menu.go` - Isolated menu model
- `tournament/` package - Tournament logic organization
- Working screen transitions with Esc to return to menu

---

### Phase 3: Single Elimination - Participant Setup 🚧 IN PROGRESS
**Status:** Planned
**Branch:** `feature/phase3-participant-setup`
**Description:** Participant configuration and dynamic bracket preview

**Goals:**
- Configure number of participants (2-64)
- Real-time bracket preview
- Dynamic bracket size calculation

**Tasks:**
1. Participant count adjustment (+/- or j/k keys)
2. Bracket preview visualization
3. Round calculation display
4. Bye position indication (for non-power-of-2 counts)
5. Validation and navigation

**Deliverables:**
- Participant setup screen
- Dynamic bracket preview
- Proceed to bracket generation

---

### Phase 4: Bracket Structure & Logic
**Status:** Planned
**Description:** Tournament bracket data structures and generation algorithms

**Goals:**
- Build tournament bracket data structure
- Calculate rounds, matches, and byes
- Handle non-power-of-2 participant counts

**Tasks:**
1. Define data structures (Match, Bracket, Player)
2. Bracket generation algorithm
3. Seeding system
4. Match tree linking

**Deliverables:**
- `tournament/bracket.go` - Core data structures
- `tournament/bracket_builder.go` - Generation algorithms
- Unit tests for bracket logic

---

### Phase 5: Bracket Visualization
**Status:** Planned
**Description:** Display full tournament bracket in TUI

**Goals:**
- ASCII art bracket rendering
- Real-time bracket updates
- Match highlighting

**Tasks:**
1. Bracket display component
2. Multi-round side-by-side layout
3. Navigation for large brackets
4. Styling with Lipgloss

**Deliverables:**
- Full bracket visualization
- Color-coded match states (pending, in-progress, completed)
- Scrollable view for large tournaments

---

### Phase 6: Match Entry & Progression
**Status:** Planned
**Description:** Record match results and progress through tournament

**Goals:**
- Enter match results
- Automatic winner advancement
- Tournament completion detection

**Tasks:**
1. Match selection and result entry
2. Winner propagation to next round
3. Tournament state tracking
4. Winner celebration screen

**Deliverables:**
- Interactive match result entry
- Automatic bracket updates
- Tournament completion flow

---

### Phase 7: Participant Naming (Enhancement)
**Status:** Planned
**Description:** Custom participant names throughout tournament

**Goals:**
- Allow custom participant names
- Display names in all views

**Tasks:**
1. Name entry screen after participant count selection
2. Default naming (Player 1, Player 2, etc.)
3. Update all views to use participant names

**Deliverables:**
- Name entry interface
- Name display in bracket
- Name editing capability

---

### Phase 8: Polish & Testing
**Status:** Planned
**Description:** Final improvements, error handling, and comprehensive testing

**Goals:**
- Improve UX
- Handle edge cases
- Test various tournament sizes

**Tasks:**
1. Error handling and validation
2. Help screens and keyboard shortcuts
3. Responsive design for various terminal sizes
4. Comprehensive testing (2, 3, 4, 8, 16, 32, 64 participants)
5. Performance optimization

**Deliverables:**
- Polished user experience
- Comprehensive error handling
- Help documentation
- Test coverage

---

### Future Phases (Post-MVP)
- **Phase 9:** Double Elimination implementation
- **Phase 10:** Round Robin implementation
- **Phase 11:** Tournament data persistence (save/load)
- **Phase 12:** Tournament history and statistics
- **Phase 13:** Export results (JSON, CSV, etc.)

---

## Implementation Priority
1. ✅ Phase 1 - Project Setup
2. ✅ Phase 2 - App Structure Refactoring
3. 🚧 Phase 3 - Participant Setup
4. ⏳ Phase 4 - Bracket Structure & Logic
5. ⏳ Phase 5 - Bracket Visualization
6. ⏳ Phase 6 - Match Entry & Progression
7. ⏳ Phase 7 - Participant Naming
8. ⏳ Phase 8 - Polish & Testing

---

## Project Structure
```
go-tournament/
├── main.go                           # Entry point, screen routing
├── screen.go                         # Screen state definitions
├── menu.go                           # Tournament selection menu
├── tournament/
│   ├── single_elimination.go         # Single elimination UI & logic
│   ├── bracket.go                    # Bracket data structures
│   └── bracket_builder.go            # Bracket generation algorithms
├── plans/
│   ├── overview.md                   # This file
│   ├── phase1-project-setup.md
│   ├── phase2-app-structure-refactoring.md
│   ├── phase3-participant-setup.md
│   └── ...
├── go.mod
├── go.sum
└── README.md
```

---

## Current Status
- **Active Branch:** `feature/phase2-app-structure-refactoring`
- **Next Milestone:** Phase 3 - Participant Setup Screen
- **Overall Progress:** ~25% (2/8 phases complete)
