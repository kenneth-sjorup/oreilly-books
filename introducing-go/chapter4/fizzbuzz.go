package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 || i%5 == 0 {
			var output string
			if i%3 == 0 {
				output = "Fizz"
			}
			if i%5 == 0 {
				output += "Buzz"
			}
			fmt.Println(i, output)
		}
	}
}
