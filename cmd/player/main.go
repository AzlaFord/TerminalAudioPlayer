package main

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/ui"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	player, err := audio.NewPlayer()
	if err != nil {
		fmt.Println("eroare new Player")
	}

	m, err := ui.NewModel(player)
	if err != nil {
		fmt.Println("erroare", err)
	}
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("erroare rulare TUI", err)
	}

}
