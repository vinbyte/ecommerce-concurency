package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var stock int = 1

var mutex sync.Mutex

func main() {

	for routine := 1; routine <= 2; routine++ {
		Wait.Add(1)
		go Routine(routine)
	}

	Wait.Wait()
	fmt.Printf("Final Stock: %d\n", stock)
}

func Routine(id int) {
	fmt.Printf("go routine %d run \n", id)
	// Lock : blocking the other goroutine to use the data
	mutex.Lock()
	if stock > 0 {
		fmt.Printf("go routine %d doing another logic process \n", id)
		time.Sleep(1 * time.Nanosecond)
		fmt.Printf("go routine %d decreasing the stock. \n", id)
		stock = stock - 1
	}
	// when we call Unlock, the other goroutine is starting to use the data
	mutex.Unlock()
	Wait.Done()
}
