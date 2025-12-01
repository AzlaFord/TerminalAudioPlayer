package main

import (
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
)

func main() {
	playlists, err := playlist.DiscoverPlaylists()
	if err != nil {
		fmt.Println("erroare", err)
		return
	}
	// dau loop pentru a returna numele la playlisturi si cantecele
	for _, playlist := range playlists {
		fmt.Println(playlist.Name)
		for _, trackname := range playlist.Tracks {
			fmt.Println(trackname.Title)
		}
	}
}
