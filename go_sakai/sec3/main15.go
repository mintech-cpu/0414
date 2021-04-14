package main

import "fmt"

// 　スライスはリサイズできる
func main() {
	n := []int{1, 2, 3}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println(board)

	n = append(n, 200, 300)
	fmt.Println(n)
}
