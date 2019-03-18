package main

import "fmt"

var cat string = "Shiro"

func main() {
	// Variables
	var x string = "Hello world"
	fmt.Println(x)

	var y string
	y = "first"
	fmt.Println(y)
	y = "second"
	fmt.Println(y)

	fmt.Println(x + " first " + y)
	fmt.Println(x == y)

	z := "Shorter statement"
	fmt.Println(z)

	// How to name a variable
	catsName := "Shiro"
	fmt.Println("My cat's name is", catsName)

	// Scope
	fmt.Println(cat)
	f()

	// Constants
	const xyz string = "Constant"

	// Defining multiple variables
	var (
		a = 1
		b = 2
		c = "three"
	)
	fmt.Println(a, b, c)

	// An example program
	example()

	// Temperature
	fahrenheitToCelsius()

	// Length
	feetToMeter()
}

func f() {
	fmt.Println(cat)
}

func example() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func fahrenheitToCelsius() {
	fmt.Print("Enter temperature in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("Temperature in celsius - ", celsius)
}

func feetToMeter() {
	fmt.Print("Feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meter := feet * 0.3048
	fmt.Println("Meter: ", meter)
}
