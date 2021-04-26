package main

import (
	"fmt"
)

func main() {
	/*
		no := 31
		if no%2 == 0 {
			fmt.Printf("%v is even of type %T\n", no, no)
		} else {
			fmt.Printf("%v is odd of type %T\n", no, no)
		}
	*/

	if no := 31; no%2 == 0 {
		fmt.Printf("%v is even of type %T\n", no, no)
	} else {
		fmt.Printf("%v is odd of type %T\n", no, no)
	}

	//for construct
	/* sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Sum = %d\n", sum) */

	//using for construct like a 'while' construct
	/*
		sum := 1
		for sum < 100 {
			sum += sum
		}
		fmt.Printf("Sum = %v\n", sum)
	*/

	//As a infinite loop
	/*
		sum := 1
		for {
			sum += sum
			if sum > 100 {
				break
			}
		}
		fmt.Printf("Sum = %v\n", sum)
	*/

	//switch case
	no := 31
	remainder := no % 2
	switch remainder {
	case 0:
		fmt.Printf("%v is even of type %T\n", no, no)
	case 1:
		fmt.Printf("%v is odd of type %T\n", no, no)
	}

	//switch case - compare against multiple values
	/*
		score
			0 - 3 -> "Terrible"
			4 - 7 -> "Not bad!"
			8 - 9 -> 'Good"
			10 -> "Excellent"
			< 0 > 10 -> "Unknown score"
	*/
	score := 13
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not bad!")
	case 8, 9:
		fmt.Println("Good!")
	case 10:
		fmt.Println("Excellent Score!")
	default:
		fmt.Println("Unknown grade!")
	}

	var n int
	fmt.Println("Enter a value")
	fmt.Scanf("%d", &n)

	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	default:
		fmt.Println("Try again")
	}

}
