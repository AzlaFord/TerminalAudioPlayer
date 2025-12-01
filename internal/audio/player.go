package audio

import (
	"bytes"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

func PlayFile(path string) {
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
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	<-readyChan
	player := otoCtx.NewPlayer(decodedMp3)
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}
}
