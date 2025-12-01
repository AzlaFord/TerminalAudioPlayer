package playlist

import (
	"log"
	"os"
)

// Deci lista de cantece in Playlist

type Track struct {
	Path  string
	Title string
}

// Lista de Playlisturi in Home/User/Music fiecare fisier in Music e socotit ca Playlist

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
		if !file.IsDir() {
			continue
		}

		var pathPlaylist = "/home/bivol/Music/" + file.Name()
		var tracksInThisPlayList []Track

		track, err := os.ReadDir(pathPlaylist)
		if err != nil {
			log.Fatal(err)
		}
		for _, trackData := range track {
			tracks := Track{
				Path:  pathPlaylist + "/" + trackData.Name(),
				Title: trackData.Name(),
			}
			tracksInThisPlayList = append(tracksInThisPlayList, tracks)
		}

		playList := Playlist{
			Name:   file.Name(),
			Tracks: tracksInThisPlayList,
		}

		playLists = append(playLists, playList)

	}

	return playLists, nil

}
