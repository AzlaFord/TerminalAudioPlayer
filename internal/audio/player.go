package player

import (
	"bytes"
	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage :player <action> <path>")
	}
	action := os.Args[1]
	path := os.Args[2]
	PlayFile(path, action)
}

func PlayFile(path, action string) {
	file, err := os.ReadFile(path)
	if err != nil {
		panic("reading" + path + "failed" + err.Error())
	}

	fileReader := bytes.NewReader(file)
	decodedMp3, err := mp3.NewDecoder(fileReader)

	if err != nil {
		panic("mp3 new decoder failed" + err.Error())
	}

	op := &oto.NewContextOptions{}
	op.SampleRate = 44100
	op.ChannelCount = 2
	op.Format = oto.FormatSignedInt16LE
	otoCtx, readyChan, err := oto.NewContext(op)
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)

	if action == "play" {
		player.Play()
	}
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
