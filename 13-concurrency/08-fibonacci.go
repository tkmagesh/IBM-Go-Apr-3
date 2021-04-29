package main

import "fmt"

func fibonacci(ch chan int) {
	count := 20
	x, y := 0, 1
	for i := 0; i < count; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	fibCh := make(chan int, 5)
	go fibonacci(fibCh)
	for fibNo := range fibCh {
		fmt.Println(fibNo)
	}
	fmt.Println("Done")
}
