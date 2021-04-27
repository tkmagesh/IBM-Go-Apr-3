package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10, 20, 30))
	fmt.Println(add(10, 20, 30, 40, 50))

	var nosSlice []int = []int{9, 4, 2, 6, 1}
	fmt.Println(add(nosSlice...))
}

func add(nos ...int) int {
	sum := 0
	for i := 0; i < len(nos); i++ {
		sum += nos[i]
	}
	return sum
}
