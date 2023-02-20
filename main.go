package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kivox/tbag/game"
	"os"
	"os/signal"
)

func main() {
	gameResp, err := tea.NewProgram(game.NewGame(), tea.WithAltScreen()).Run()
	if err != nil {
		panic(err)
	}

	if gameResp.(*game.Game).Dead {
		fmt.Println("Thanks for playing, you lost!")
	} else if gameResp.(*game.Game).Won {
		fmt.Println("Thanks for playing, you won!")
	} else {
		fmt.Println("Thanks for playing, you quit the game!")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
}
