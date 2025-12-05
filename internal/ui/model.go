package ui

import (
	"TerminalAudioPlayer/internal/playlist"
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	title, desc string
	index       int
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
	return nil
}

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
	var items []list.Item

	if err != nil {
		return Model{}, err
	}
	// incarca playlisturile in playlistItem

	for idx, pl := range listPl {
		items = append(items, item{title: pl.Name, desc: " ", index: idx})
	}

	// incarca playlisturile in trackList
	for _, pl := range tracks {
		fmt.Println(pl.Title)
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Playlists"

	if len(listPl) > 0 {
		tracks = listPl[0].Tracks
	}

	return Model{
		playlists:       listPl,
		tracks:          tracks,
		status:          "ready",
		focusOnPlaylist: true,
		playListItem:    l,
	}, nil
}
