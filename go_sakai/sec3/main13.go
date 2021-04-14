package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int型からfloat型への型変換
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f", xx, xx, xx)
	// float型からint型への型変換
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d", yy, yy, yy)

	// string型からint型への型変換
	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d", i, i, i)

}
