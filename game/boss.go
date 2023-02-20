package game

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"math"
)

type BossKeyMap struct {
	Fight key.Binding
}

func (b BossKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{b.Fight}
}

func (b BossKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{b.Fight},
	}
}

var bossKeys = BossKeyMap{
	Fight: key.NewBinding(
		key.WithKeys("f", "enter"),
		key.WithHelp("f/enter", "fight"),
	),
}

type BossFight struct {
	game *Game

	keys BossKeyMap
	help help.Model

	health int

	move           int
	currentMessage string
	bossHealth     int
}

func (b *BossFight) Init() tea.Cmd {
	return nil
}

func (b *BossFight) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.keys.Fight):
			if b.move >= 10 {
				b.currentMessage = "You outlasted the boss!"
				return b, nil
			} else {
				if b.bossHealth <= 0 {
					b.game.currentMessage = "You defeated the boss!"
					return b.game, nil
				} else if b.health <= 0 {
					b.game.currentMessage = "You died!"
					b.game.Dead = true
					return b.game, nil
				} else {
					bossDamage := int(math.Ceil(1*((float64(b.game.prompts)/10.0)+1)) * (float64(b.game.swordLevel) + 1) * ((float64(b.game.strengthMultiplier) / 10.0) + 1))
					charDamage := int(math.Floor(1 * ((float64(b.game.prompts) / 10.0) + 1)))

					b.health -= charDamage
					b.bossHealth -= bossDamage

					b.currentMessage = fmt.Sprintf("You dealt %v damage to the boss, the boss struct back and dealt %v damage!", bossDamage, charDamage)

					b.move++

					return b, nil
				}
			}
		}
	}

	return b, nil
}

func (b *BossFight) View() string {
	var response string

	response += fmt.Sprintf("Prompt: %d\n", b.game.prompts)
	response += fmt.Sprintf("Move: %d\n", b.move)
	response += "---\n"
	response += fmt.Sprintf("Health: %d\n", b.health)
	response += fmt.Sprintf("Attack Power: %f ♥\n", (1+(float64(b.game.swordLevel)/10.0))*(b.game.strengthMultiplier+1))
	response += "---\n"
	response += fmt.Sprintf("Boss Health: %d\n", b.bossHealth)

	response += "\n\n\n"

	response += b.currentMessage

	response += "\n\n\n"

	response += b.help.View(b.keys)

	return response
}
