package main

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
)

func main() {

	var index int
	var index2 int
	var cmd string

	playlists, err := playlist.DiscoverPlaylists()
	if err != nil {
		fmt.Println("erroare", err)
	}
	audio.Init()
	for {
		fmt.Scanln(&cmd)
		selectedPlaylists := playlists[index]

		switch cmd {

		case "k":
			for i, playlist := range playlists {
				fmt.Println(i, playlist.Name)
			}
		case "l":
			for i, song := range selectedPlaylists.Tracks {
				fmt.Println(i, song.Title)
			}
		case "s":
			fmt.Println("Introdu index playlist")
			fmt.Scanln(&index)

		case "t":
			fmt.Println("Introdu index cantec")
			fmt.Scanln(&index2)
		case "p":
			selectSong := selectedPlaylists.Tracks[index2]
			err = audio.PlayFile(selectSong.Path)
			if err != nil {
				fmt.Println("Eroare la redare", err)
			}
		}

		if cmd == "q" {
			break
		}
	}
}
