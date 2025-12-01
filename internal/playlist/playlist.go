package playlist

import (
	"os"
	"path/filepath"
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

func DiscoverPlaylists() ([]Playlist, error) {

	home, erruser := os.UserHomeDir()

	if erruser != nil {
		return nil, erruser
	}

	var playLists []Playlist

	musicDir := filepath.Join(home, "Music")
	files, err := os.ReadDir(musicDir)

	if err != nil {
		return nil, err
	}

	// da loop in directori si gaseste alte directori/ playlisturi
	for _, file := range files {
		// verifica daca e Directory
		if !file.IsDir() {
			continue
		}
		// facut o varaibila locala si un path pentru a fi putea folosit in structura playlist
		pathPlaylist := filepath.Join(musicDir, file.Name())
		var tracksInThisPlayList []Track

		track, err := os.ReadDir(pathPlaylist)

		if err != nil {
			return nil, err
		}
		// cauta catece in playlisturi/directoriu
		for _, trackData := range track {

			if trackData.IsDir() {
				continue
			}
			pathTrackData := filepath.Join(pathPlaylist, trackData.Name())
			var extension = filepath.Ext(trackData.Name())
			switch extension {
			case ".mp3", ".wav", ".webm", ".ogg":
				tracks := Track{
					Path:  pathTrackData,
					Title: trackData.Name(),
				}
				tracksInThisPlayList = append(tracksInThisPlayList, tracks)
			default:

			}

		}

		playList := Playlist{
			Name:   file.Name(),
			Tracks: tracksInThisPlayList,
		}

		playLists = append(playLists, playList)

	}

	return playLists, nil

}
