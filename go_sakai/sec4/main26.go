package main

import "fmt"

func main() {
	l := []string{"H", "I", "S"}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	m := map[string]int{"H": 1, "I": 2, "S": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
