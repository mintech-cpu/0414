package main

import "fmt"

func main() {
	var i int = 1 // 関数外でも宣言できる
	var f64 float64 = 1.2
	var s1, s2 string = "hello", "go"
	var t, f bool = true, false
	fmt.Println(i, f64, s1, s2, t, f)
	fmt.Printf("%T", s1)

	xi := 1 // 関数ないでしか宣言できない
	xi = 2
	fmt.Println(xi)
	fmt.Printf("%T", xi)

}
