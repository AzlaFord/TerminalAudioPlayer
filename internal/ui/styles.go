package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

var volume = lipgloss.NewStyle().Background(lipgloss.Color("91"))

// stilurile la lista cu playlisturi
var docStyle = lipgloss.NewStyle().MarginTop(1).BorderStyle(lipgloss.NormalBorder())

func (m Model) View() string {
	player := m.player

	m.help = help.New()
	m.help.ShowAll = true

	if !m.focusOnPlaylist {
		// daca e focus pe false se va aplica situruile la lista playlisturi si tabel
		docStyle = docStyle.BorderForeground(lipgloss.Color("240"))
		baseStyle = baseStyle.BorderForeground(lipgloss.Color("111"))
	} else {
		docStyle = docStyle.BorderForeground(lipgloss.Color("111"))
		baseStyle = baseStyle.BorderForeground(lipgloss.Color("240"))
	}

	tableMusic := baseStyle.Render(m.table.View()) + "\n"
	style := lipgloss.NewStyle().Background(lipgloss.Color("91"))
	styled := style.Render("TUI Music Player")
	// de rezolvat problmea crash cand dau volume - cand nu e inizitalizat playerul cred
	volumeStilat := volume.Render("Volume :" + strconv.Itoa(int(player.Volume*100)))
	volume := lipgloss.Place(30, 10, lipgloss.Center, lipgloss.Bottom, volumeStilat)
	block := lipgloss.Place(30, 10, lipgloss.Center, lipgloss.Bottom, styled)

	blocTitle := lipgloss.JoinHorizontal(0, block, volume)

	list := lipgloss.Place(30, 10, lipgloss.Center, lipgloss.Bottom, tableMusic)
	// aici am combinat lista cu playlisuri si tabelul folosind JoinHorizontal
	final := lipgloss.JoinHorizontal(0.05, docStyle.Render(m.playListItem.View()), list)

	return blocTitle + final
}
