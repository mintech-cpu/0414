package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "W", "w", 1))
	fmt.Println(s) // コピーしているだけなので、元の値は変化なし
	// s = strings.Replace(s, "W", "w", 1)　置き換えたい場合は、sに代入する
	fmt.Println(strings.Contains(s, "Hello"))

	fmt.Println("test" +
		"test")

}
