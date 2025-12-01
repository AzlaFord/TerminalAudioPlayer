package main

import (
	"TerminalAudioPlayer/internal/audio"
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
)

func main() {
	playlists, err := playlist.DiscoverPlaylists()
	if err != nil {
		fmt.Println("erroare", err)
		return
	}

	pl := playlists[0]
	t := pl.Tracks[0]
	fmt.Println("Piesa", t.Title)
	audio.PlayFile(t.Path)

	// dau loop pentru a returna numele la playlisturi si cantecele
	for _, playlist := range playlists {
		fmt.Println(playlist.Name)
		for _, trackname := range playlist.Tracks {
			fmt.Println(trackname.Title)
		}
	}
}
