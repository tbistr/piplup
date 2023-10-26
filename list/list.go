package list

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
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
	cursor        int

	keymap KeyMap
}

func New(items []Item, width, height int) Model {
	return Model{
		items:   items,
		Width:   width,
		Height:  height,
		Spacing: 1,
		keymap:  DefaultKeyMap(),
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.CursorUp):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keymap.CursorDown):
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		}
	}
	var cmd tea.Cmd
	return m, cmd
}

func (m Model) View() string {
	items := []string{}
	currentHeight := 0

	for i, item := range m.items {
		renderred := item.Render(i == m.cursor)
		currentHeight += item.Height() + m.Spacing

		if currentHeight > m.Height {
			break
		}

		items = append(items, renderred)
	}

	return strings.Join(items, strings.Repeat("\n", m.Spacing+1))
}
