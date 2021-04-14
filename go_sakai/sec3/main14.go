package main

import (
	"fmt"
)

// 配列はリサイズできない
func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	c := [...]int{100, 200}
	fmt.Println(c)
}
