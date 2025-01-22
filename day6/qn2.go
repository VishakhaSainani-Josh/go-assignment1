/*2. Given a string, reverse it using one go routine.And inside go routine print reversed string and number of goroutines launched

e.g: Input: test123 output: 321tset 2*/

package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func reverse(s string) {
	var reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversed = reversed + string(s[i])
	}

	fmt.Print(reversed," ")
	fmt.Print(runtime.NumGoroutine())

}
func main() {
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()
		reverse(input)
	}()

	wg.Wait()

}
