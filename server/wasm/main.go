package main

import "syscall/js"

func foo(i []js.Value) {
	js.Global().Set("output", js.ValueOf(i[0].String()))
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("foo", js.NewCallback(foo))
	<-c
}
