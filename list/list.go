package list

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Item interface {
	Render(selected bool) string
	Height() int
}

type Model struct {
	// Items is the list of items to display.
	items         []Item
	Width, Height int
	Spacing       int
}

func New(items []Item, width, height int) Model {
	return Model{
		items:   items,
		Width:   width,
		Height:  height,
		Spacing: 1,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() string {
	items := []string{}
	currentHeight := 0

	for _, item := range m.items {
		renderred := item.Render(false)
		currentHeight += item.Height() + m.Spacing

		if currentHeight > m.Height {
			break
		}

		items = append(items, renderred)
	}

	return strings.Join(items, strings.Repeat("\n", m.Spacing+1))
}
