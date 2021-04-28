package main

import (
	"fmt"
	"time"
)

func print(message string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(delay * time.Second)
	}
}

func main() {
	go print("Hello", 2)
	go print("World", 1)
	var input string
	fmt.Scanln(&input)
}
