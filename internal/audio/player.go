package audio

import (
	"bytes"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

type Player struct {
	otoCtx          *oto.Context
	currentPlayer   *oto.Player
	Volume          float64
	pause           bool
	prevValueVolume float64
}

func NewPlayer() (*Player, error) {

	p := &Player{Volume: 1.0}

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
// volume e practic rezultatul la default volume + sau - la step
func (p *Player) SetVolume(volume float64) error {
	if volume > 1.0 {
		volume = 1.0
	}

	if volume < 0.0 {
		volume = 0
	}

	p.Volume = volume
	if p.currentPlayer == nil {
		p.Volume = volume
		return nil
	}
	p.currentPlayer.SetVolume(p.Volume)
	return nil
}

func (p *Player) SetMute() bool {

	if p.Volume != 0.0 {
		p.prevValueVolume = p.GetVolume()
		p.SetVolume(0.0)
		return true
	}
	if p.Volume == 0.0 {
		p.SetVolume(p.prevValueVolume)
		return false
	}
	return false
}

func (p *Player) PlayFile(path string) error {
	p.pause = false
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
	p.currentPlayer.SetVolume(p.Volume)
	p.currentPlayer.Play()

	return nil
}

func (p *Player) IsPlaying() bool {
	return p.currentPlayer.IsPlaying()
}

func (p *Player) TogglePlayPause() {

	if p.currentPlayer == nil {
		return
	}
	if p.currentPlayer.IsPlaying() {
		p.currentPlayer.Pause()
		p.pause = true
	} else {
		p.currentPlayer.Play()
		p.pause = false
	}
}

func (p *Player) IsPaused() bool {
	return p.pause
}
