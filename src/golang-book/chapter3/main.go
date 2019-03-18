package main

import (
	"fmt"
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
	fmt.Println("true && true =", true && true)
	fmt.Println("true && false =", true && false)
	fmt.Println("true || true =", true || true)
	fmt.Println("true || false =", true || false)
	fmt.Println(!true)
	var sum = 32132 * 42452
	fmt.Println(sum, reflect.TypeOf(sum))
	var bool = (true && false) || (false && true) || !(false && true)
	fmt.Println(bool)
}
