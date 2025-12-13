package ui

import (
	"TerminalAudioPlayer/internal/playlist"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TrackStartingMsg struct {
	Title string
}

type TrackErrorMsg struct {
	Err error
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	player := m.player

	switch msg := msg.(type) {
	case TickMsg:
		if m.shouldAutoNext() {
			m.selectedTrack++
			m.table.SetCursor(m.selectedTrack)
			return m, tea.Batch(m.playTrackCmd(m.tracks[m.selectedTrack]), tickCmd())
		}
		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.focusOnPlaylist = true
			m.table.Blur()
		case "q":
			return m, tea.Quit
		case "enter":
			m.focusOnPlaylist = false
			m.table.Focus()
		case "=":
			player.IncreaseVolume(0.05)
		case " ":
			player.TogglePlayPause()
		case "m":
			player.SetMute()
		case "b":
			if m.canHitPrev() {
				m.selectedTrack--
				m.table.SetCursor(m.selectedTrack)
				return m, m.playTrackCmd(m.tracks[m.selectedTrack])
			}
		case "n":
			if m.canHitNext() {
				m.selectedTrack++
				m.table.SetCursor(m.selectedTrack)
				return m, m.playTrackCmd(m.tracks[m.selectedTrack])
			}
		case "-":
			player.DecreaseVolume(0.05)
		case "r":
			if len(m.tracks) == 0 {
				break
			}
			if !m.focusOnPlaylist {
				m.selectedTrack = m.table.Cursor()
				idx := m.table.Cursor()
				if idx >= 0 && idx < len(m.tracks) {
					tracks := m.tracks[idx]
					return m, tea.Batch(m.playTrackCmd(tracks), tickCmd())
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

func (m Model) shouldAutoNext() bool {
	player := m.player

	if player == nil {
		return false
	}
	if player.IsPaused() {
		return false
	}
	if player.IsPlaying() {
		return false
	}
	if m.selectedTrack+1 >= len(m.tracks) {
		return false
	}
	return true

}

func (m Model) canHitNext() bool {
	p := m.player

	if p == nil {
		return false
	}
	if m.selectedTrack+1 >= len(m.tracks) {
		return false
	}
	return true
}

func (m Model) canHitPrev() bool {
	p := m.player

	if p == nil {
		return false
	}
	if m.selectedTrack-1 < 0 {
		return false
	}
	return true
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

func tickCmd() tea.Cmd {
	return tea.Tick(200*time.Millisecond, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}
