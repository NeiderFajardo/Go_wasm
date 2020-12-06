package main

import(
	"strconv"
	"syscall/js"
	"log"
)


func sumar(this js.Value, i []js.Value) interface {}{
	valor1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	valor2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	a,_ := strconv.Atoi(valor1)
	b,_ := strconv.Atoi(valor2)
	log.Println(a, b)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", a+b)
	
	return a+b

}

func restar(this js.Value, i []js.Value) interface {}{
	valor1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	valor2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	a,_ := strconv.Atoi(valor1)
	b,_ := strconv.Atoi(valor2)
	log.Println(a, b)
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", a-b)
	return a-b
}

func registerCallbacks() {
	js.Global().Set("sumar", js.FuncOf(sumar))
	js.Global().Set("restar", js.FuncOf(restar))
}

func main(){
	c := make(chan struct{}, 0)

	println("WASM Go Inicializado")

	registerCallbacks()
	<-c
}