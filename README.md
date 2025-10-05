# Go Tournament Manager

A terminal-based tournament management application built with Go and Bubbletea.

## Features

- **Beautiful TUI Interface**: Built using [Bubbletea](https://github.com/charmbracelet/bubbletea) for an interactive terminal experience
- **Multiple Tournament Types**: Support for different tournament formats
- **Elegant Card-based Navigation**: Navigate through tournament options using intuitive card interface

## Tournament Types

### 🥊 Single Elimination
The classic tournament format where participants are eliminated after losing a single match. Winners advance through rounds until only one remains.

### 🔄 Double Elimination (Coming Soon)
Players get a second chance in the loser's bracket, making tournaments more forgiving and exciting.

### 🔁 Round Robin (Coming Soon)
Every participant plays against every other participant at least once, ensuring fair competition.

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-tournament
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run main.go
```

## Controls

- **← → or h l**: Navigate between tournament type cards
- **Enter**: Select tournament type
- **q or Ctrl+C**: Quit the application

## Requirements

- Go 1.24.0 or later
- Terminal with Unicode support for best experience

## Dependencies

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal User Interface framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout for terminal applications

## Development Status

This project is currently in active development. The tournament selection screen is complete, and the Single Elimination tournament implementation is in progress.

## Contributing

Contributions are welcome! Please feel free to submit issues and enhancement requests.