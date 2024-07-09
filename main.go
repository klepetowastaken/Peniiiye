package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	for i := 0; i < 10; i++ {
		{
			fmt.Printf("Jsem super presne tolikrat: %v\n", i+1)

			fmt.Println("zdravím, World!")
		}
	}

}
