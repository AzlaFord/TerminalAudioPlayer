package ui

import (
	"TerminalAudioPlayer/internal/playlist"

	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	playlists        []playlist.Playlist
	selectedPlaylist int
	playListItem     list.Model
	trackList        list.Model

	tracks        []playlist.Track
	selectedTrack int

	status          string
	focusOnPlaylist bool
}

func NewModel() (Model, error) {
	listPl, err := playlist.DiscoverPlaylists()
	var tracks []playlist.Track

	if err != nil {
		return Model{}, err
	}

	if len(listPl) > 0 {
		tracks = listPl[0].Tracks
	}

	return Model{
		playlists:       listPl,
		tracks:          tracks,
		status:          "ready",
		focusOnPlaylist: true,
	}, nil
}
