package audio

import (
	"bytes"
	"errors"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

var (
	otoCtx        *oto.Context
	currentPlayer *oto.Player
)

func Init() error {

	if otoCtx != nil {
		return nil
	}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100

	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return err
	}
	<-readyChan
	otoCtx = ctx
	return nil

}

func SetVolume(volume float64) error {

	if currentPlayer == nil {
		return errors.New("nu exista playerul")
	}
	if volume >= 0 && volume <= 100 {
		currentPlayer.SetVolume(volume / 100)
	}
	if volume > 100 {
		currentPlayer.SetVolume(1)
	}
	if volume < 0 {
		currentPlayer.SetVolume(0)
	}
	return nil
}

func PlayFile(path string) error {

	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileReader := bytes.NewReader(file)
	decodedMp3, err := mp3.NewDecoder(fileReader)

	if err != nil {
		return err
	}

	if currentPlayer != nil {
		currentPlayer = nil
	}
	player := otoCtx.NewPlayer(decodedMp3)
	currentPlayer = player

	currentPlayer.Play()

	return nil
}
