package main

import "fmt"

func init() {
	fmt.Println("Init!") // メイン関数より先に表示される
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello World")
}
