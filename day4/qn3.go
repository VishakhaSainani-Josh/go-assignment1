/*3. The given program takes an integer value as input from the user i.e 1 or 2.
Option 1 represents Rectangle and Option 2 represents Square.

Given the metrics of the shapes are hard-coded, complete the given program to achieve the following:

1. Create an interface Quadrilateral which has the following method signatures
Area() int
Perimeter() int

2. Implement the Quadrilateral interface for the given shapes i.e Rectangle and Square

3. Define a "Print" function which accepts any shape that implements Quadrilateral interface and Prints Area and Perimeter of shape in the following manner:

"Area :  <value>"
"Perimeter :  <value>"


HINT: Step 2 means, to define methods in Quadrilateral interface for given shapes


Formulae:

Area of Rectangle: length * width
Perimeter of Rectangle: 2*(length + breadth)


Area of Square: side * side
Perimeter of Square: 4 * side*/

package main

import (
	"fmt"
	"log"
)

type Quadrilateral interface {
	Area() int
	Perimeter() int
}

type rectangle struct {
	length  int
	breadth int
}

type square struct {
	side int
}

func (r rectangle) Area() int {
	return r.length * r.breadth
}

func (r rectangle) Perimeter() int {
	return 2 * (r.length + r.breadth)
}

func (s square) Area() int {
	return s.side * s.side
}

func (s square) Perimeter() int {
	return 4 * s.side
}

func printResult(shape Quadrilateral) {
	fmt.Printf("Area : %d\n", shape.Area())
	fmt.Printf("Perimeter : %d", shape.Perimeter())
}

func main() {

	var input int

	if _, err := fmt.Scan(&input); err != nil {
		log.Fatal(err)
	}

	switch input {
	case 1:
		rect := rectangle{length: 10, breadth: 20}
		printResult(rect)
	case 2:
		sq := square{side: 10}
		printResult(sq)
	}

}
