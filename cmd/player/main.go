package main

import (
	"TerminalAudioPlayer/internal/playlist"
	"fmt"
)

func main() {
	playlist, err := playlist.DiscoverPaylists()
	if err != nil {
		fmt.Println("erroare", err)
		return
	}
	fmt.Println(playlist)
}
