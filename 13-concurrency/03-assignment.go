package main

import (
	"fmt"
	"time"
)

func print(message string, delay time.Duration, ch1 chan string, ch2 chan string) {
	for i := 0; i < 5; i++ {
		<-ch1
		fmt.Println(message)
		time.Sleep(delay * time.Second)
		ch2 <- "done"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", 2, ch1, ch2)
	go print("World", 1, ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
	fmt.Println("Job Done!")
}
