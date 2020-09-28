package test

import (
	"os"
	"testing"

	mp3 "github.com/hajimehoshi/go-mp3"
	"github.com/triole/oto"
)

func TestAud(t *testing.T) {
	f, err := os.Open("../tfl.mp3")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	de, err := mp3.NewDecoder(f)
	if err != nil {
		t.Error("mpe decode: ", err)
	}

	ctx, _ := oto.NewContext(de.SampleRate(), 2, 2, 2048)

	player := ctx.NewPlayer()
	bytes := make([]byte, de.Length())
	_, e := de.Read(bytes)
	if e != nil {
		t.Error(e)
	}
	player.Write(bytes)

}
