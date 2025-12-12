package ui

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/key"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	CursorUp    key.Binding
	CursorDown  key.Binding
	NextPage    key.Binding
	PrevPage    key.Binding
	GoToStart   key.Binding
	GoToEnd     key.Binding
	Filter      key.Binding
	ClearFilter key.Binding

	CancelWhileFiltering key.Binding
	AcceptWhileFiltering key.Binding

	ShowFullHelp  key.Binding
	CloseFullHelp key.Binding

	Quit      key.Binding
	ForceQuit key.Binding
}

type Model struct {
	playlists        []playlist.Playlist
	keyMap           KeyMap
	selectedPlaylist int
	playListItem     list.Model
	trackList        list.Model
	table            table.Model
	tracks           []playlist.Track
	selectedTrack    int

	player          *audio.Player
	status          string
	focusOnPlaylist bool
	percent         float64
	progress        progress.Model
	mute            bool
}

type item struct {
	title, desc string
	index       int
}

var barStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
var baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Background(lipgloss.Color("240"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

const (
	padding  = 2
	maxWidth = 80
)
const listHeight = 20

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
	columns := []table.Column{{Title: "Order", Width: 10},
		{Title: "Title", Width: 35},
		{Title: "Playlist", Width: 15},
		{Title: "Duration", Width: 9},
	}

	var rows []table.Row
	for i, song := range tracks {
		idx := strconv.Itoa(i + 1)
		rows = append(rows, table.Row{idx, song.Title, song.PlaylistTitle, "3"})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(false),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true).Bold(false)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(false)
	t.SetStyles(s)

	return t
}

func NewModel(player *audio.Player) (Model, error) {

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
	}

	const defaultWidth = 20
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Playlists"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// verific daca lista exista macar si iau primul element adica lista cu cantece
	if len(listPl) > 0 {
		tracks = listPl[0].Tracks
	}

	tbl := NewTable(tracks)

	return Model{
		playlists:       listPl,
		tracks:          tracks,
		status:          "ready",
		focusOnPlaylist: true,
		playListItem:    l,
		table:           tbl,
		player:          player,
		mute:            false,
		keyMap:          DefaultKeyMap(),
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

func DefaultKeyMap() KeyMap {
	return KeyMap{
		// Browsing.
		CursorUp: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "up"),
		),
		CursorDown: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "down"),
		),
		PrevPage: key.NewBinding(
			key.WithKeys("left", "h", "pgup", "b", "u"),
			key.WithHelp("←/h/pgup", "prev page"),
		),
		NextPage: key.NewBinding(
			key.WithKeys("right", "l", "pgdown", "f", "d"),
			key.WithHelp("→/l/pgdn", "next page"),
		),
		GoToStart: key.NewBinding(
			key.WithKeys("home"),
			key.WithHelp("g/home", "go to start"),
		),
		GoToEnd: key.NewBinding(
			key.WithKeys("end"),
			key.WithHelp("G/end", "go to end"),
		),
		Filter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "filter"),
		),
		ClearFilter: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "clear filter"),
		),

		// Filtering.
		CancelWhileFiltering: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "cancel"),
		),
		AcceptWhileFiltering: key.NewBinding(
			key.WithKeys("e", "shift+tab", "ctrl+k", "up", "ctrl+j", "down"),
			key.WithHelp("e", "apply filter"),
		),

		// Toggle help.
		ShowFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "more"),
		),
		CloseFullHelp: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "close help"),
		),

		// Quitting.
		Quit: key.NewBinding(
			key.WithKeys("q", "esc"),
			key.WithHelp("q", "quit"),
		),
		ForceQuit: key.NewBinding(key.WithKeys("ctrl+c")),
	}
}
