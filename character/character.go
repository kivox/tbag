package character

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Selection  int
	Strength   int
	Luck       int
	PointCount int
}

type Option struct {
	label string
	value int
}

var options = []Option{
	{"Strength", 0},
	{"Luck", 0},
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "esc", "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.Selection--
			if m.Selection < 0 {
				m.Selection = len(options) - 1
			}
		case "down":
			m.Selection++
			if m.Selection >= len(options) {
				m.Selection = 0
			}
		case "left":
			if options[m.Selection].value > 0 {
				options[m.Selection].value--
				m.PointCount++
			}
		case "right":
			if m.PointCount > 0 {
				options[m.Selection].value++
				m.PointCount--
			}
		}

		// Ensure that the total point count is always 5 or less
		total := 0
		for _, o := range options {
			total += o.value
		}
		if total > 5 {
			if options[m.Selection].value > 0 {
				options[m.Selection].value--
				m.PointCount++
			} else {
				for i, o := range options {
					if o.value > 0 {
						options[i].value--
						m.PointCount++
						break
					}
				}
			}
		}

		// Update the model
		m.Strength = options[0].value
		m.Luck = options[1].value

		return m, nil
	default:
		return m, nil
	}
}

func (m Model) View() string {
	var menu string
	for i, o := range options {
		var symbol string
		if i == m.Selection {
			symbol = "* "
		} else {
			symbol = "  "
		}
		menu += fmt.Sprintf("%s%s: %d\n", symbol, o.label, o.value)
	}

	status := fmt.Sprintf("Points remaining: %d", m.PointCount)

	return fmt.Sprintf("%s\n\n%s", menu, status)
}
