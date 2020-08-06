package main

import (
	"syscall/js"
)

func fib(this js.Value,params []js.Value) interface{} {
	n := params[0].Int()
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

 
// 将Go里面的方法注入到window.fibNative里面
func registerCallbacks() {
	js.Global().Set("fibNative", js.FuncOf(fib))
}

func main() {
	registerCallbacks()
	select {}
}