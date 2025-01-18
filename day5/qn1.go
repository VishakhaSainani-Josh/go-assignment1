/*1. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value at that index in slice>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.

Complete the given code to recover from panic inside the accessSlice function, when index out of range panic occurs.
Also, Print the following after handling panic

"internal error: <recovered panic value>"


Example Test Case 1 :
Input: 3
Output:
item 3, value 6
Testing Panic and Recover

*/

package main

import (
	"fmt"
	"log"
)

func accessSlice(s1 []int, index int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Internal error", r)
		}
	}()

	fmt.Printf("Item %d,value %d", index, s1[index])
	fmt.Println("\nTesting Panic and Recover")
}

func main() {
	s1 := []int{6, 5, 4, 3, 2, 1}

	var input int
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	accessSlice(s1, input)

}
