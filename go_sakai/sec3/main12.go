package main

import (
	"fmt"
)

func main() {
	var t, f bool = true, false
	fmt.Printf("%T %v %t", t, t, t)
	fmt.Printf("%T %v %t", f, f, f)

}
