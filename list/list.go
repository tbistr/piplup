package list

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Item interface {
	Render() string
}

type Model struct {
	// Items is the list of items to display.
	items []Item
}

func New(items []Item) Model {
	return Model{
		items: items,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() string {
	b := strings.Builder{}

	for _, item := range m.items {
		b.WriteString(item.Render())
		b.WriteString("\n")
	}

	return b.String()
}
