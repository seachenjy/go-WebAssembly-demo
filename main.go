package main

import (
	"io/ioutil"
	"net/http"
	"syscall/js"
)

func fib(this js.Value, params []js.Value) interface{} {
	n := params[0].Int()
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func makebyte(this js.Value, params []js.Value) interface{} {
	go func() {
		res, err := http.Get("xxxxxxxx")
		if err == nil {
			if bytes, err := ioutil.ReadAll(res.Body); err == nil {
				// dst := js.Global().Get("Uint8Array").New(len(bytes))
				// js.CopyBytesToJS(dst, bytes)
				js.Global().Get("document").Call("write", string(bytes))
			}
		}
	}()
	return true
}

func play(this js.Value, params []js.Value) interface{} {
	go func() {
		c2dx := js.Global().Get("document").Call("getElementById", "main").Call("getContext", "2d")
		c2dx.Call("beginPath")
		c2dx.Set("fillStyle", "#dd0000")
		c2dx.Call("arc", 100, 100, 16, 0, 3.1415*2)
		c2dx.Call("fill")
		c2dx.Call("closePath")
	}()
	return true
}

// 将Go里面的方法注入到window.fibNative里面
func registerCallbacks() {
	js.Global().Set("fibNative", js.FuncOf(fib))
	js.Global().Set("makebyte", js.FuncOf(makebyte))
	js.Global().Set("play", js.FuncOf(play))
}

func main() {
	registerCallbacks()
	select {}
}
