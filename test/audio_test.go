package test

import (
	"io"
	"net/http"
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

func TestLoader(t *testing.T) {
	response, err := http.Get("https://svs.gsfc.nasa.gov/vis/a000000/a004700/a004720/lroc_color_poles_4k.tif")
	if err != nil {
		t.Error(err)
	}
	defer response.Body.Close()
	buf := make([]byte, 1024)
	file, err := os.Create("./4k.tif")
	defer file.Close()
	if err != nil {
		t.Error(err)
	} else {
		for {

			n, err := response.Body.Read(buf)
			if err != nil && err != io.EOF {
				break
			}

			wn, werr := file.Write(buf)
			if werr != nil {
				break
			}
			t.Log(n, wn)
		}
	}
	t.Log("done")
}
