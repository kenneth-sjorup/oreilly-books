package main

import "fmt"

func main () {
	fmt.Println(
		`
******************* FEET -> METERS ******************
********************* CONVERTER *********************
             `)

	// read user input
	fmt.Print("Enter Number of Feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	// convert to meters
	output := input * .3048

	fmt.Println("Number of Meters: ", output)
}
