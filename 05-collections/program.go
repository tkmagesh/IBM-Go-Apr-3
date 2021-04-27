package main

import "fmt"

func main() {
	//Array
	var nos [5]int
	for i := 0; i < 5; i++ {
		nos[i] = i
	}
	fmt.Println("nos length = ", len(nos))
	fmt.Println(nos)

	//initilization while creation
	//var products [5]string = [5]string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	var products = []string{"Pen", "Pencil", "Marker", "Scribble Pad", "Mouse"}
	fmt.Printf("Type of products = %T\n", products)
	fmt.Printf("No of products = %v\n", len(products))

	//Iterating using range
	for _, value := range products {
		fmt.Println(value)
	}

	//multidimenstional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row-%d*col-%d", i, j)
		}
	}
	fmt.Printf("Matrix - %v\n", matrix)
}
