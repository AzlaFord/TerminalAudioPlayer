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
		return docStyle.Render(m.trackList.View())
	}

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
		case "e":
			m.focusOnPlaylist = true
		case "q":
			return m, tea.Quit
		case "enter":
			m.focusOnPlaylist = false
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.playListItem.SetSize(msg.Width-h, msg.Height-v)
	}
	var cmd tea.Cmd
	m.playListItem, cmd = m.playListItem.Update(msg)
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
