package main

import "fmt"

// クロージャー

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	func(x int) {
		x++
		fmt.Println(x)
	}(1)

}
