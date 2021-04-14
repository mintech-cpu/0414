package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)
	fmt.Println(result)

	if result == "ok" {
		fmt.Println("great")
	}

	// num := 6
	// if num%2 == 0 {
	// 	fmt.Println("2")
	// } else if num%3 == 0 {
	// 	fmt.Println("3")
	// } else {
	// 	fmt.Println("else")
	// }

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
}
