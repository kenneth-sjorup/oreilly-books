package main

import "fmt"

func main() {
	fmt.Println(
		`
*************** FAHRENHEIT -> CELSIUS ***************
********************** CONVERTER ********************
             `)

	// read user input
	fmt.Print("Enter Temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	// convert to celsius
	output := (input - 32) * 5/9

	fmt.Println( "Temperature in Celsius: ", output)
}
