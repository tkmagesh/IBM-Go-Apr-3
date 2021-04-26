package main

import "fmt"

func main() {
	for {
		choice := getMenuOption()
		if choice == 5 {
			break
		}
		result := 0
		n1, n2 := getOperands()
		switch choice {
		case 1:
			result = add(n1, n2)
		case 2:
			result = subtract(n1, n2)
		case 3:
			result = multiply(n1, n2)
		case 4:
			result = divide(n1, n2)
		default:
			fmt.Println("Invalid choice. Please try again!")
		}
		fmt.Printf("Result = %d\n", result)
	}
}

func add(n1, n2 int) int {
	return n1 + n2
}

func subtract(n1, n2 int) int {
	return n1 - n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func divide(n1, n2 int) int {
	return n1 / n2
}

func getOperands() (n1 int, n2 int) {
	fmt.Println("Enter Operand 1")
	fmt.Scanf("%d", &n1)
	fmt.Println("Enter Operand 2")
	fmt.Scanf("%d", &n2)
	return
}

func getMenuOption() int {
	var choice int
	fmt.Println("Enter your choice:")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Scanf("%d", &choice)
	return choice
}
