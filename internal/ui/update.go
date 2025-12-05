package ui

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"

	// "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1)

type item struct {
	title, desc string
}

func (i item) Title() string {
	return i.title
}

func (i item) Description() string {
	return i.desc
}

func (i item) FilterValue() string {
	return i.title
}

func (m Model) Init() tea.Cmd {
	plist , _ := playlist.DiscoverPlaylists()

	var items [] m.playListItem.Item
	
	for _,pl := range plist{
		items = append (items,m.playlists{title:pl.Name})
	}

	m.playListItem.SetItems(items)
	
	return nil
}

func (m Model) View() string {
	if m.focusOnPlaylist {
		return docStyle.Render(m.playListItem.View())
	}
	if !m.focusOnPlaylist {
		return docStyle.Render(m.trackList.View())
	}

	return "ceva"
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

			if m.selectedPlaylist > 0 && m.focusOnPlaylist {
				m.selectedPlaylist--
			}
		case "down":
			if m.selectedPlaylist < len(m.playlists)-1 && m.focusOnPlaylist {
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
		case "esc":
			m.focusOnPlaylist = true
		case "q":
			return m, tea.Quit
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
	var cmd tea.Cmd
	m.playlistList, cmd = m.playlistList.Update(msg)
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
