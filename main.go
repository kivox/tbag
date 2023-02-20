package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kivox/tbag/character"
)

func main() {
	charConfig, err := tea.NewProgram(character.Model{
		Selection:  0,
		Strength:   0,
		Luck:       0,
		PointCount: 5,
	}).Run()
	if err != nil {
		panic(err)
	}

	_ = charConfig.(character.Model)
}
