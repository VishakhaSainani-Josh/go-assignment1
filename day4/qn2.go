/*2. The given program accepts 2 values from the user, length and breadth of a rectangle respectively.

Complete the program by writing methods "Area" and "Perimeter" on Rectangle type to calculate the area and perimeter of a rectangle.

Hint:
Method Signatures for Area and Perimeter
Area() int
Perimeter() int

Formulae:
Area = length * width
Perimeter = 2 * (length + width)

Example Test Case 1:
Input: 10 20
Output:
Area of Rectangle: 200
Perimeter of Rectangle: 60*/

package main

import (
	"fmt"
	"log"
)

type rectangle struct {
	length  int
	breadth int
}

func (r rectangle) Area() int {
	return r.length * r.breadth
}

func (r rectangle) Perimeter() int {
	return 2 * (r.length + r.breadth)
}

func main() {

	var length int
	var breadth int

	if _, err := fmt.Scanf("%d %d", &length, &breadth); err != nil {
		log.Fatal(err)
	}

	rect := rectangle{length: length, breadth: breadth}
	fmt.Printf("Area of rectangle: %d\n", rect.Area())
	fmt.Printf("Perimeter of rectangle: %d", rect.Perimeter())

}
