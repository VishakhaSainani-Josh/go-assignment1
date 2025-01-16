/*1. The given program accepts an integer value between 1 to 4 from the user.
There is an option associated with each value, which is basically objects of different data types with some associated value.

Write a method named "AcceptAnything" which takes any type of data as input.

Based on the option chosen by the user, we will send different types of objects to this function and it should print the following based on its respective data type of value sent to the function "AcceptAnything".

integer :
"This is a value of type Integer, <value>"

string :
"This is a value of type String, <value>"

boolean :
"This is a value of type Boolean, <value>"

Custom data type Hello :
"This is a value of type Hello, <value>"*/

package main

import (
	"fmt"
	"log"
)

type Hello float32

type data interface{}

func AcceptAnything(d data) {
	switch d.(type) {
	case int:
		fmt.Printf("This is a value of type Integer, %d", d)
	case string:
		fmt.Printf("This is a value of type String, %s", d)
	case bool:
		fmt.Printf("This is a value of type Boolean, %t", d)
	case Hello:
		fmt.Printf("This is a value of type Hello, %v", d)

	}
	//In above cases we can do fmt.Printf("This is a value of type %T,%v",d,d) but here solved because of output given in the question

}

func main() {
	var input int
	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	switch input {
	case 1:
		AcceptAnything(1)
	case 2:
		AcceptAnything("Hello, Good Morning")
	case 3:
		AcceptAnything(false)
	case 4:
		var h Hello
		AcceptAnything(h)
	}

}
