package main

import "fmt"

func foo(params ...int) {

}

func main() {
	// Q1. 以下の1.11をint型に変換して出力してください。
	f := 1.11
	ff := int(f)
	fmt.Printf("%T %v %d", ff, ff, ff)

	// Q2. コードを書かずに以下の出力結果を答えてください。
	// s := []int{1, 2, 5, 6, 2, 3, 1}
	// fmt.Println(s[2:4])
	fmt.Println(5, 6)

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)

}
