/*Assignment:
In the below code snippet concurrent goroutines execution corrupts a piece of data by
accessing it simultaneously it leads in raise condition.
https://go.dev/play/p/SmLMf8hZQzs
output when you run this : 1 is Even (may vary when you run multiple times)

Update the given code to print the correct output.
output once code is corrected: 0 is Even*/

//The follwoing code was given
//Updates made: Added MUtex to acquire lock to critical section

package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

var mutex sync.Mutex

func main() {

	n := 0
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		nIsEven := isEven(n)

		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")

	}()
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
