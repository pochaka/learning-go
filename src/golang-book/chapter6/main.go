package main

import "fmt"

func main() {
	x := [5]float64{98, 93, 77, 82, 83}
	var total float64
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
	sliceAppend()
	sliceCopy()
	maps()
	periodicTable()
	smallestNumber()
}

func sliceAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func maps() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)

	y := make(map[int]int)
	y[1] = 10
	y[0] = 1000
	delete(y, 1)
	fmt.Println(y)
}

func periodicTable() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Berylium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	fmt.Println(elements["Li"])
	fmt.Println(elements["in"])
	name, ok := elements["in"]
	fmt.Println(name, ok)
	if name, ok := elements["in"]; ok {
		fmt.Println(name, ok)
	}
}

func smallestNumber() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	fmt.Println(len(x))
	min := x[0]
	for _, value := range x {
		fmt.Println(value, min)
		if min > value {
			min = value
		}
	}
	fmt.Println(min)
}
