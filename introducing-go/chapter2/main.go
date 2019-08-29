package main

import "fmt"

func main() {
	// working with numeric literals
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)

	// working with strings
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	// working with booleans
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// exercises
	fmt.Println(32132 * 42542)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
