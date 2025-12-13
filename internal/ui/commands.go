package ui

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	LineUp       key.Binding
	LineDown     key.Binding
	PageUp       key.Binding
	PageDown     key.Binding
	HalfPageUp   key.Binding
	HalfPageDown key.Binding
	GotoTop      key.Binding
	GotoBottom   key.Binding
}
type KeyMapList struct {
	Down   key.Binding
	Up     key.Binding
	Back   key.Binding
	Select key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		LineUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "sus"),
		),
		LineDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "jos"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("b/pgup", "page up"),
		),
		PageDown: key.NewBinding(
			key.WithKeys("f", "pgdown"),
			key.WithHelp("f/pgdn", "page down"),
		),
	}
}

func ListDefaultKeyMap() KeyMapList {
	return KeyMapList{
		Down:   key.NewBinding(key.WithKeys("j", "down", "ctrl+n"), key.WithHelp("j", "jos")),
		Up:     key.NewBinding(key.WithKeys("k", "up", "ctrl+p"), key.WithHelp("k", "up")),
		Back:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "esc")),
		Select: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	}
}

func (k KeyMapList) ShortHelp() []key.Binding {
	return []key.Binding{k.Back, k.Select}
}

func (k KeyMapList) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Back, k.Select},
	}
}
