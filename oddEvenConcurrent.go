package main

import (
	"fmt"
	"sync"
	"time"
)

func printEven(n int, op chan int, ip chan int, wg *sync.WaitGroup) {
	for i := 0; i < n; i += 2 {
		fmt.Print(i, " ")
		time.Sleep(2000)
		ip <- 0
		<-op
	}
	wg.Done()
}

func printOdd(n int, op chan int, ip chan int, wg *sync.WaitGroup) {
	for i := 1; i < n; i += 2 {
		<-op
		fmt.Print(i, " ")
		ip <- 1
	}
	wg.Done()
}

func main() {
	N := 10

	op := make(chan int)
	ip := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go printEven(N, op, ip, &wg)
	go printOdd(N, ip, op, &wg)

	wg.Wait()
}
