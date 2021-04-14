package main

import (
	"fmt"
	"time"
)

func osName() string {
	return "windows"
}

func main() {
	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac", os)
	case "windows":
		fmt.Println("win", os)
	default:
		fmt.Println("default", os)
	}

	switch os := osName(); os {
	case "mac":
		fmt.Println("mac", os)
	case "windows":
		fmt.Println("win", os)
	default:
		fmt.Println("default", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch { // switchの後ろの条件式は書かなくても処理可能
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")

	}

}
