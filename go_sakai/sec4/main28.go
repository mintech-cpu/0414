package main

import "fmt"

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	foo()

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("foo")

	// file, _ := os.Open("lesson")
	// defer file.Close()
	// data := make([]byte, 100)
	// file.Read(data)
	// fmt.Println(string(data))

}
