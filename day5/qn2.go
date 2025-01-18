/*2. In the given code, the accessSlice function accepts a slice and index.
If the value is present in the slice at that index, the program should print the following.

"item <index>, value <value present at the index>"

But if the index does not hold any value,
it will lead to an index out of range panic in our program.


So in order to safeguard our program from panicking, add a condition to handle the condition if

index > lengthOfSlice - 1

and return an error from the accessSlice function if the above condition is met.
The error message should be the following :

"length of the slice should be more than index"

Complete the given program to return an error from the accessSlice function and handle the returned error inside the main function to print the message.

Example Test Case 1 :
Input: 3
Output:
item 3, value 6*/

package main

import (
	"errors"
	"fmt"
	"log"
)

func accessSlice(s []int, index int) (int, error) {
	if index<0{
		return 0, errors.New("index should not be negative")

	}
	if len(s)-1 < index   {
		return 0, errors.New("length of slice should be more than index")
	}
	return s[index], nil

}

func main() {

	var slice = []int{6, 5, 4, 3, 2, 1}
	var input int

	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	val, err := accessSlice(slice, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Item %d,value %d", input, val)

}
