package main

import ("fmt"
	"reflect"
)

func main() {
	var a = 1 + 1
	var b = 1.0 + 1.0
	var c = (len("Hello world"))
	var d = "Hello world"[1]
	var e = "Hello" + "World"
	fmt.Println(a, reflect.TypeOf(a))	
	fmt.Println(b, reflect.TypeOf(b))
	fmt.Println(c, reflect.TypeOf(c))
	fmt.Println(d, reflect.TypeOf(d))	
	fmt.Println(e, reflect.TypeOf(e))
	
}
