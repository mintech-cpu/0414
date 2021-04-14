package main

import "fmt"

func add(x, y int) (int, int) {
	return x * y, x + y
}

func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(1, 3)
	fmt.Println(r1, r2)

	r3 := cal(100, 200)
	fmt.Println(r3)

	f := func(x int) {
		fmt.Println("inner", x)
	}
	f(2)

	func(x int) {
		fmt.Println("inner", x)
	}(2)
}
