package main

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/ui"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	audio.Init()

	m, err := ui.NewModel()
	if err != nil {
		fmt.Println("erroare", err)
	}
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("erroare rulare TUI", err)
	}

}
