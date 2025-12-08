package audio

import (
	"bytes"
	"errors"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Player struct {
	otoCtx        *oto.Context
	currentPlayer *oto.Player
	Volume        float64
}

func NewPlayer() (*Player, error) {

	p := &Player{Volume: 100.0}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100

	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	ctx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return nil, err
	}
	<-readyChan
	p.otoCtx = ctx
	return p, nil

}

// ia volumul
func (p *Player) GetVolume() float64 {
	return p.Volume
}

// cresterea volumului
func (p *Player) IncreaseVolume(step float64) {
	p.SetVolume(p.GetVolume() + step)
}

// scaderea volumului
func (p *Player) DecreaseVolume(step float64) {
	p.SetVolume(p.GetVolume() - step)
}

// seteaza volumul apelat in DecreseVolume si IncreaseVolume
func (p *Player) SetVolume(volume float64) error {

	if p.currentPlayer == nil {
		return errors.New("nu exista playerul")
	}

	if volume > 100 {
		volume = 100
	}

	if volume < 0 {
		volume = 0
	}

	p.Volume = volume / 100
	p.currentPlayer.SetVolume(p.Volume)
	return nil
}

func (p *Player) PlayFile(path string) error {

	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fileReader := bytes.NewReader(file)
	decodedMp3, err := mp3.NewDecoder(fileReader)

	if err != nil {
		return err
	}

	if p.currentPlayer != nil {
		p.currentPlayer = nil
	}
	player := p.otoCtx.NewPlayer(decodedMp3)
	p.currentPlayer = player

	p.currentPlayer.Play()

	return nil
}
