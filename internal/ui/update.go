package ui

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	var b strings.Builder

	fmt.Println(&b)

	for i, pl := range m.playlists {
		prefix := " "
		if m.focusOnPlaylist && i == m.selectedPlaylist {
			prefix = "> "
		}
		fmt.Fprintf(&b, "%s%s\n", prefix, pl.Name)
	}

	fmt.Fprintln(&b, "")
	fmt.Fprintln(&b, "Tracks :")

	for i, t := range m.tracks {
		prefix := " "
		if !m.focusOnPlaylist && i == m.selectedTrack {
			prefix = "> "
		}
		fmt.Fprintf(&b, "%s%s\n", prefix, t.Title)
	}

	return b.String()

}

type TrackStartingMsg struct {
	Title string
}
type TrackErrorMsg struct {
	Err error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.selectedPlaylist > 0 {
				m.selectedPlaylist--
			}
		case "down":
			if m.selectedPlaylist < len(m.playlists)-1 {
				m.selectedPlaylist++
			}
		case "j":
			if m.selectedTrack > 0 && m.focusOnPlaylist == false {
				m.selectedTrack--
			}
		case "k":
			if m.selectedTrack < len(m.tracks)-1 && m.focusOnPlaylist == false {
				m.selectedTrack++
			}
		case "q":
			break
		case "enter":
			if m.focusOnPlaylist {
				if len(m.playlists) == 0 {
					break
				}
				m.tracks = m.playlists[m.selectedPlaylist].Tracks
				m.selectedTrack = 0
				m.focusOnPlaylist = false
			} else {
				if len(m.tracks) == 0 {
					break
				}
				track := m.tracks[m.selectedTrack]
				return m, playTrackCmd(track)
			}

		}

	}
	return m, nil
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
