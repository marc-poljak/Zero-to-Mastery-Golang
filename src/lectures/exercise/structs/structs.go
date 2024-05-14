//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

func area(r Rectangle) float64 {
	return r.length * r.width
}

func perimeter(r Rectangle) float64 {
	return (r.length + r.width) * 2
}

func printResults(r Rectangle) {
	fmt.Println("Area:", area(r))
	fmt.Println("Perimeter:", perimeter(r))
}

func main() {
	r := Rectangle{length: 10, width: 5}
	printResults(r)

	r.length *= 2
	r.width *= 2
	printResults(r)

}
