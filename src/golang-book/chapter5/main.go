package main

import "fmt"

func main(){
	i := 1
	for i <= 10 {
		switch i {
		case 0: fmt.Print("Zero, ")
		case 1: fmt.Print("One, ")
		case 2: fmt.Print("Two, ")
		case 3: fmt.Print("Three, ")
		case 4: fmt.Print("Four, ")
		case 5: fmt.Print("Five, ")
		case 6: fmt.Print("Six, ")
		case 7: fmt.Print("Seven, ")
		case 8: fmt.Print("Eight, ")
		case 9: fmt.Print("Nine, ")
		default: fmt.Print("Unknown number, ")
		}
		if i % 2 == 0{
			fmt.Println("even,", i)
		} else {
			fmt.Println("odd,", i)
		}
		i += 1
	}
	divisible_by_three()
}

func divisible_by_three(){
	counter := 1
	for i := 1; i <= 100; i++{
		if i % 3 == 0{
			fmt.Println(i, "is divisible by 3")
			counter += 1
		}	
	}
	fmt.Println("Total amount of divisible by 3 numbers between 1 and 100 is,", counter)
}


