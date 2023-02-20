package game

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kivox/tbag/character"
	"math"
	"math/rand"
)

type GameKeyMap struct {
	Shop    key.Binding
	Forward key.Binding
	Quit    key.Binding
}

func (k GameKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Shop, k.Forward, k.Quit},
	}
}

var gameKeys = GameKeyMap{
	Shop: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "open shop"),
	),
	Forward: key.NewBinding(
		key.WithKeys("f", "w", "enter"),
		key.WithHelp("enter", "move forward"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func (k GameKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Shop, k.Forward, k.Quit}
}

type Game struct {
	help help.Model
	keys GameKeyMap

	coins       int
	swordLevel  int
	armourLevel int

	luckMultiplier     float64
	strengthMultiplier float64

	prompts        int
	currentMessage string

	bossFights int

	Dead bool
	Won  bool
}

func (g *Game) Init() tea.Cmd {
	return nil
}

func NewGame() *Game {
	charConfig, _ := tea.NewProgram(character.Model{
		Selection:  0,
		Strength:   0,
		Luck:       0,
		PointCount: 5,
	}).Run()

	charModel := charConfig.(character.Model)

	return &Game{
		keys: gameKeys,
		help: help.New(),

		coins: 0,

		swordLevel:  0,
		armourLevel: 0,

		luckMultiplier:     (float64(charModel.Luck) / 10.0) + 1,
		strengthMultiplier: (float64(charModel.Strength) / 10.0) + 1,

		prompts:        0,
		currentMessage: "You wake up in a dark room. You don't know where you are, but you know you can only move forward.",

		bossFights: 0,

		Dead: false,
		Won:  false,
	}
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if g.Dead {
		return g, tea.Quit
	}

	if g.bossFights >= 5 {
		g.Won = true
		return g, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, g.keys.Shop):
			return &Shop{
				keys: shopKeys,
				help: help.New(),
				game: g,
			}, nil
		case key.Matches(msg, g.keys.Forward):
			g.prompts++

			randNum := rand.Intn(100)
			if randNum > 75 {
				g.bossFights++
				return &BossFight{
					game:           g,
					keys:           bossKeys,
					help:           help.New(),
					health:         int(math.Ceil(10 * ((float64(g.armourLevel) / 10.0) + 1))),
					move:           0,
					currentMessage: "You encounter a boss! You must defeat it to move forward.",
					bossHealth:     int(math.Ceil(1.0 + ((float64(g.prompts) / 10.0) + 1))),
				}, nil
			} else {
				coins := math.Ceil(float64(rand.Intn(10)) * g.luckMultiplier)
				g.currentMessage = fmt.Sprintf("You moved forward.\nYou found a pot of gold!\nYou gained %v coins.", coins)
				g.coins += int(coins)
				return g, nil
			}
		case key.Matches(msg, g.keys.Quit):
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g *Game) View() string {
	var response string

	response += fmt.Sprintf("Coins: %d\n", g.coins)
	response += "---\n"
	response += fmt.Sprintf("Luck Multiplier: %f\n", g.luckMultiplier)
	response += fmt.Sprintf("Strength Multiplier: %f\n", g.strengthMultiplier)
	response += "---\n"
	response += fmt.Sprintf("Sword Level: %d\n", g.swordLevel)
	response += fmt.Sprintf("Armour Level: %d\n", g.armourLevel)
	response += "---\n"
	response += fmt.Sprintf("Prompts: %d\n", g.prompts)

	response += "\n\n\n"

	response += g.currentMessage

	response += "\n\n\n"

	response += g.help.View(g.keys)

	return response
}
