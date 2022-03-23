package main

import (
	"sync"
)

var (
	concurrent     = 1000000
	totalExecution = 2
)

func main() {
	UseArray()
}

func Synchronous() int {
	var (
		deposit = 0
	)

	for i := 1; i <= totalExecution; i++ {
		func() {
			for i := 1; i <= concurrent; i++ {
				deposit = deposit + 1
			}
		}()
	}

	return deposit
}

func UseChannel() int {
	var (
		deposit = 0
		total   = make(chan int)
		wg      = new(sync.WaitGroup)
	)
	wg.Add(totalExecution)

	for i := 1; i <= totalExecution; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 1; i <= concurrent; i++ {
				deposit = deposit + 1
			}
			total <- deposit
		}(wg)
		<-total
	}

	wg.Wait()
	return deposit
}

func UseMutex() int {
	var (
		mtx     sync.Mutex
		deposit = 0
		wg      = new(sync.WaitGroup)
	)
	wg.Add(totalExecution)

	for i := 1; i <= totalExecution; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 1; i <= concurrent; i++ {
				mtx.Lock()
				deposit = deposit + 1
				mtx.Unlock()
			}
		}(wg)
	}

	wg.Wait()
	return deposit
}

func UseArray() int {
	var (
		temp [2]int
		sum  int
		wg   = new(sync.WaitGroup)
	)

	wg.Add(totalExecution)
	for i := 0; i < totalExecution; i++ {
		go func(wg *sync.WaitGroup, index int) {
			defer wg.Done()
			for i := 1; i <= concurrent; i++ {
				temp[index] = temp[index] + 1
			}
		}(wg, i)
	}
	wg.Wait()

	for i := 0; i < totalExecution; i++ {
		sum += temp[i]
	}

	return sum
}
