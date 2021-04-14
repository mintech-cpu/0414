package main

import "fmt"

const Pi = 3.14 // constの値は書き換え不可

const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)

}
