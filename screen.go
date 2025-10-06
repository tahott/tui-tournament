package main

// Screen represents different screens/states in the application
type Screen int

const (
	ScreenMenu Screen = iota
	ScreenSingleElimination
	ScreenDoubleElimination
	ScreenRoundRobin
)
