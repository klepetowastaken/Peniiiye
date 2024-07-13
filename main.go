package main

import "fmt"

/*
	func Funkce(x int, y int) int {
		return x / y
	}
*/

func Vypinac(x, y *bool) {
	if *x == true {
		*x = false
	} else {
		*x = true
	}

	if *y == true {
		*y = false
	} else {

		*y = true
	}

}

func main() {

	var x bool = false
	var y bool = false

	Vypinac(&x, &y)

	fmt.Println(x, y)

}
