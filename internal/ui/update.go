package ui

import (
	"TerminalAudioPlayer/internal/playlist"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("111")).MarginTop(1).BorderStyle(lipgloss.ASCIIBorder())

func (m Model) View() string {

	tableMusic := baseStyle.Render(m.table.View()) + "\n"
	style := lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(11)).Background(lipgloss.Color("91"))
	styled := style.Render("TUI Music Player")
	block := lipgloss.Place(30, 10, lipgloss.Center, lipgloss.Bottom, styled)
	list := lipgloss.Place(30, 10, lipgloss.Center, lipgloss.Bottom, tableMusic)
	final := lipgloss.JoinHorizontal(0.05, docStyle.Render(m.playListItem.View()), list)

	return block + final

}

type TrackStartingMsg struct {
	Title string
}

type TrackErrorMsg struct {
	Err error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	step := m.player

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// de rezolvat problema esc da crash
			m.focusOnPlaylist = true
			m.table.Blur()
		case "q":
			return m, tea.Quit
		case "enter":
			m.focusOnPlaylist = false
			m.table.Focus()
		case "=":
			step.IncreaseVolume(0.05)
		case "space":
			// de rezolvat problema cu space la pauza
			step.Pause()
		case "-":
			step.DecreaseVolume(0.05)
		case "r":
			if len(m.tracks) == 0 {
				break
			}
			if !m.focusOnPlaylist {
				idx := m.table.Cursor()
				if idx >= 0 && idx < len(m.tracks) {
					tracks := m.tracks[idx]
					return m, m.playTrackCmd(tracks)
				}

			}

		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.playListItem.SetSize(msg.Width-h, msg.Height-v)
	}

	if m.focusOnPlaylist {
		m.playListItem, cmd = m.playListItem.Update(msg)
		idx := m.playListItem.Index()
		if idx != m.selectedPlaylist {
			m.selectedPlaylist = idx
			if idx >= 0 && idx < len(m.playlists) {
				m.tracks = m.playlists[idx].Tracks
				m.table = NewTable(m.tracks)
			}
		}
	} else {
		m.table, cmd = m.table.Update(msg)
	}
	return m, cmd
}

func (m Model) playTrackCmd(t playlist.Track) tea.Cmd {

	p := m.player

	return func() tea.Msg {
		err := p.PlayFile(t.Path)
		if err != nil {
			return TrackErrorMsg{Err: err}
		}
		return TrackStartingMsg{Title: t.Title}
	}
}
