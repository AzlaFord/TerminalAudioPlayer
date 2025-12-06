package ui

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1)

func (m Model) View() string {

	if m.focusOnPlaylist {
		return docStyle.Render(m.playListItem.View())
	} else {
		return baseStyle.Render(m.table.View()) + "\n"
	}
}

type TrackStartingMsg struct {
	Title string
}

type TrackErrorMsg struct {
	Err error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	volume := 100.0
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "e":
			m.focusOnPlaylist = true
			m.table.Blur()
		case "q":
			return m, tea.Quit
		case "enter":
			m.focusOnPlaylist = false
			m.table.Focus()
		case "=":
			audio.SetVolume(volume)
		case "-":
			audio.SetVolume(volume)
		case "r":
			if len(m.tracks) == 0 {
				break
			}
			if !m.focusOnPlaylist {
				idx := m.table.Cursor()
				if idx >= 0 && idx < len(m.tracks) {
					tracks := m.tracks[idx]
					return m, playTrackCmd(tracks)
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

func playTrackCmd(t playlist.Track) tea.Cmd {
	return func() tea.Msg {
		err := audio.PlayFile(t.Path)
		if err != nil {
			return TrackErrorMsg{Err: err}
		}
		return TrackStartingMsg{Title: t.Title}
	}
}
