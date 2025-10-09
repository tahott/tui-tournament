# Phase 2: Application Structure Refactoring

**Status:** ✅ COMPLETED
**Branch:** `feature/phase2-app-structure-refactoring`
**Date Completed:** 2025-10-06

## Overview
Refactor the monolithic application structure into a modular, scalable architecture that supports multiple screens and tournament types. This phase establishes the foundation for adding tournament-specific functionality.

## Goals
- Separate concerns: menu screen vs tournament screens
- Implement screen state management system
- Create modular architecture for different views
- Enable seamless navigation between screens

## Tasks Completed

### 1. Screen State Enum Creation ✅
**File:** `screen.go`

Created a type-safe screen state system:
```go
type Screen int

const (
    ScreenMenu Screen = iota
    ScreenSingleElimination
    ScreenDoubleElimination
    ScreenRoundRobin
)
```

This allows the application to track which screen is currently active and route accordingly.

### 2. Menu Extraction ✅
**File:** `menu.go`

Extracted all menu-related code from `main.go` into a dedicated module:
- `menuModel` struct with tournament selection logic
- `tournamentType` struct with screen mapping
- All menu-specific styling (cardStyle, selectedCardStyle, etc.)
- Navigation and selection handling
- Custom message type `screenChangeMsg` for screen transitions

**Key Addition:** Screen change messaging
```go
type screenChangeMsg struct {
    screen Screen
}
```

When user presses Enter, the menu sends a `screenChangeMsg` to trigger screen transition.

### 3. Tournament Package Structure ✅
**Directory:** `tournament/`
**File:** `tournament/single_elimination.go`

Created dedicated package for tournament logic:
- `SingleEliminationModel` struct
- Placeholder UI showing "Coming soon..."
- Independent Update/View methods
- Tournament-specific styling

This establishes the pattern for adding other tournament types.

### 4. Main Model Refactoring ✅
**File:** `main.go`

Completely restructured the main model:

**New Model Structure:**
```go
type model struct {
    currentScreen     Screen
    menuModel         menuModel
    singleElimination tournament.SingleEliminationModel
    width             int
    height            int
}
```

**Key Features:**
- Screen state tracking (`currentScreen`)
- Delegation pattern - routes messages to appropriate sub-model
- Screen switching logic
- ESC key handler to return to menu from any screen

**Update Method:**
```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Handle global keys (quit, esc)
    // Handle screen change messages
    // Delegate to current screen's model
}
```

**View Method:**
```go
func (m model) View() string {
    switch m.currentScreen {
    case ScreenMenu:
        return m.menuModel.View()
    case ScreenSingleElimination:
        return m.singleElimination.View()
    // ...
    }
}
```

## Deliverables

### Files Created/Modified
```
go-tournament/
├── main.go                           # ✨ Refactored - screen routing
├── screen.go                         # ✨ New - screen state enum
├── menu.go                           # ✨ New - menu module
├── tournament/
│   └── single_elimination.go         # ✨ New - SE tournament screen
├── go.mod
├── go.sum
└── README.md
```

### Key Features
- ✅ Screen state management system
- ✅ Modular architecture with clear separation of concerns
- ✅ Working screen transitions (Menu → Single Elimination)
- ✅ ESC key to return to menu
- ✅ Foundation for adding more tournament types

## Architecture Pattern

### Delegation Pattern
Each screen has its own model with `Update()` and `View()` methods. The main model:
1. Receives all messages
2. Handles global actions (quit, esc)
3. Handles screen transitions (screenChangeMsg)
4. Delegates other messages to the active screen's model
5. Renders the active screen's view

### Message Flow
```
User Input → Main Update() → Global handling
                           → Screen-specific delegation
                           → Sub-model Update()
                           → Return updated model + commands
```

### View Rendering
```
Main View() → Check currentScreen
           → Call appropriate sub-model's View()
           → Return rendered string
```

## Technical Details

### Screen Transitions
1. User selects tournament in menu
2. Menu model returns `screenChangeMsg`
3. Main model receives message
4. Updates `currentScreen` field
5. Next render uses new screen's View()

### Navigation Back to Menu
1. User presses ESC
2. Main model's Update() catches it
3. Sets `currentScreen = ScreenMenu`
4. Menu is rendered on next frame

## Benefits of This Architecture
- ✅ **Modularity:** Each screen is independent
- ✅ **Scalability:** Easy to add new tournament types
- ✅ **Maintainability:** Clear file organization
- ✅ **Testability:** Can test each screen in isolation
- ✅ **Flexibility:** Screens can have different models/state

## Testing
Manually tested:
- ✅ Application builds without errors
- ✅ Menu screen displays correctly
- ✅ Navigation between tournament cards works
- ✅ Selecting Single Elimination transitions to new screen
- ✅ ESC returns to menu from Single Elimination screen
- ✅ Window resize handled correctly in both screens
- ✅ Quit (q/Ctrl+C) works from any screen

## Lessons Learned
- Bubbletea's message-passing architecture works well with delegation patterns
- Type-safe enums (iota) are excellent for screen state management
- Separating concerns early makes future development much easier
- Custom message types (`screenChangeMsg`) are powerful for inter-module communication

## Next Steps
→ **Phase 3:** Single Elimination - Participant Setup
- Implement participant count selection (2-64)
- Dynamic bracket preview
- Real-time bracket calculation
- Proceed to bracket generation
