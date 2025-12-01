package playlist

import (
	"log"
	"os"
)

type Track struct {
	Path  string
	Title string
}

type Playlist struct {
	Name   string
	Tracks []Track
}

func DiscoverPaylists() ([]Playlist, error) {

	var playLists []Playlist

	files, err := os.ReadDir("/home/bivol/Music")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		playList := Playlist{
			Name:   file.Name(),
			Tracks: []Track{},
		}
		playLists = append(playLists, playList)
	}
	return playLists, nil

}
