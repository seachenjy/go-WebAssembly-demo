package test

import (
	"fmt"
	"io/ioutil"

	"github.com/tosone/minimp3"
	"github.com/triole/oto"
)

func test() {
	var file, _ = ioutil.ReadFile("../tfl.mp3")
	dec, data, _ := minimp3.DecodeFull(file)

	ctx, _ := oto.NewContext(dec.SampleRate, dec.Channels, 2, 2048)

	fmt.Println("rate", dec.SampleRate)
	fmt.Println("kbps", dec.Kbps)
	fmt.Println("channels", dec.Channels)
	player := ctx.NewPlayer()
	player.Write(data)
}
