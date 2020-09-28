package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://get.flashfxp.com/ftp/client/download/FlashFXP54_3970_Setup.exe")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	buf := make([]byte, 1024)
	f, err := os.Create("./FlashFXP54_3970_Setup.exe")
	defer f.Close()
	if err != nil {
		panic(err)
	} else {
		for {

			n, err := response.Body.Read(buf)
			if err != nil && err != io.EOF {
				break
			}

			if err == io.EOF {
				break
			}
			wn, _ := f.Write(buf[:n])
			fmt.Println(len(buf), wn)
		}
	}
	fmt.Println("done")
}
