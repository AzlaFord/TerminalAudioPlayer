package ui

import (
	"TerminalAudioPlayer/internal/playlist"
)

type Model struct {
	playlists        []playlist.Playlist
	selectedPlaylist int

	tracks        []playlist.Track
	selectedTrack int

	status string
}

func NewModel() (Model, error) {
	list, err := playlist.DiscoverPlaylists()
	var tracks []playlist.Track

	if err != nil {
		return Model{}, err
	}

	if len(list) > 0 {
		tracks = list[0].Tracks
	}

	return Model{
		playlists:        list,
		selectedPlaylist: 0,
		tracks:           tracks,
		selectedTrack:    0,
		status:           "ready",
	}, nil
}
