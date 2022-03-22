package main

import (
	"sync"
)

var (
	total      = make(chan int)
	deposit    = 0
	concurrent = 1000000
)

func topup() {
	deposit = deposit + 1
}

func execute(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= concurrent; i++ {
		topup()
	}
	total <- deposit
}
func main() {
	totalExecution := 2
	wg := new(sync.WaitGroup)
	wg.Add(totalExecution)

	for i := 1; i <= totalExecution; i++ {
		go execute(wg)
		<-total
	}

	wg.Wait()
}
