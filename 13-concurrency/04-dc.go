package main

import (
	"fmt"
	"time"
)

func sumSync(nos []int) int {
	result := 0
	for _, no := range nos {
		result += no
	}
	return result
}

func sum(nos []int, resultCh chan int) {
	result := 0
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}

func main() {

	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	firstResultCh := make(chan int)
	secondResultCh := make(chan int)
	firstSet := data[:len(data)/2]
	secondSet := data[len(data)/2:]

	start := time.Now()
	go sum(firstSet, firstResultCh)
	go sum(secondSet, secondResultCh)
	total := <-firstResultCh + <-secondResultCh
	fmt.Println(total)
	elapsed := time.Since(start)
	fmt.Printf("Time taken (async) : %v\n", elapsed)

	start = time.Now()
	syncTotal := sumSync(firstSet) + sumSync(secondSet)
	fmt.Println(syncTotal)
	elapsed = time.Since(start)
	fmt.Printf("Time taken (sync) : %v\n", elapsed)
	/* total :=  sum(firstSet) + sum(secondSet)
	fmt.Println(total) */
	/*
		1. break the data into 2 equal sets
		2. run 2 goroutine instances of sum function one for each set
		3. get the result from each of them and aggregate and print the final result
	*/
}
