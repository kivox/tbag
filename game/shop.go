package game

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ShopKeyMap struct {
	Back key.Binding
}

func (k ShopKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Back},
	}
}

var shopKeys = ShopKeyMap{
	Back: key.NewBinding(
		key.WithKeys("esc", "q", "ctrl+c"),
		key.WithHelp("q", "leave shop"),
	),
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k ShopKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Back}
}

type Shop struct {
	game *Game
	keys ShopKeyMap
	help help.Model
}

func (s *Shop) Init() tea.Cmd {
	return nil
}

func (s *Shop) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.keys.Back):
			return s.game, nil
		}
	}

	return s, nil
}

func (s *Shop) View() string {
	response := "The shop is currently closed, come back later!\n\n"
	response += s.help.View(s.keys)
	return response
}
