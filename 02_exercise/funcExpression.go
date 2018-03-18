package main

import "fmt"

func main() {
	half := func(val int) (int, bool) {
		return val / 2, val%2 == 0
	}
	fmt.Println(half(1))
	fmt.Println(half(2))
}
