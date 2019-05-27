package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println("Average", average(xs))
	fmt.Println("Add", add(1, 2, 3))
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(int(xs[1]), int(xs[2])))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	defer second()
	first()
}

func average(xs []float64) float64 {
	total := 0.0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
