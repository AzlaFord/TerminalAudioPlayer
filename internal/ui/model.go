package ui

import (
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	playlists        []playlist.Playlist
	selectedPlaylist int
	playListItem     list.Model
	trackList        list.Model
	table            table.Model
	tracks           []playlist.Track
	selectedTrack    int

	status          string
	focusOnPlaylist bool
}

type item struct {
	title, desc string
	index       int
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

const listHeight = 14

type itemDelegate struct{}

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

func NewTable(tracks []playlist.Track) table.Model {
	columns := []table.Column{
		{Title: "Order", Width: 4},
		{Title: "Title", Width: 12},
		{Title: "Location", Width: 10},
	}

	var rows []table.Row
	for i, song := range tracks {
		idx := strconv.Itoa(i + 1)
		rows = append(rows, table.Row{idx, song.Title, song.Path})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

func NewModel() (Model, error) {

	listPl, err := playlist.DiscoverPlaylists()
	var tracks []playlist.Track
	var items []list.Item

	if err != nil {
		return Model{}, err
	}

	// incarca playlisturile in playlistItem
	for _, pl := range listPl {
		length := "Tracks " + strconv.Itoa(len(pl.Tracks))
		items = append(items, item{title: pl.Name, desc: length})
		tracks = append(tracks, pl.Tracks...)
	}

	const defaultWidth = 20
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Playlists"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	tbl := NewTable(tracks)

	if len(listPl) > 0 {
		tracks = listPl[0].Tracks
	}

	return Model{
		playlists:       listPl,
		tracks:          tracks,
		status:          "ready",
		focusOnPlaylist: true,
		playListItem:    l,
		table:           tbl,
	}, nil
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s %s ", index+1, i.title, i.desc)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
