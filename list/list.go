package list

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item interface {
	Render(selected bool) string
	Height() int
}

type Model struct {
	// Items is the list of items to display.
	items         []Item
	Width, Height int
}

func New(items []Item, width, height int) Model {
	return Model{
		items:  items,
		Width:  width,
		Height: height,
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
		currentHeight += item.Height()

		if currentHeight > m.Height {
			break
		}

		items = append(items, renderred)
	}

	return lipgloss.JoinVertical(lipgloss.Top, items...)
}
