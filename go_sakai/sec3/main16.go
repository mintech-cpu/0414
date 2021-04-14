package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v", len(n), cap(n), n)

	a := make([]int, 3) // 引数を一つにするとlen, capともに3になる
	fmt.Printf("len=%d cap=%d value=%v", len(a), cap(a), a)

	c := make([]int, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	d := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		d = append(d, i)
		fmt.Println(d)
	}

}
