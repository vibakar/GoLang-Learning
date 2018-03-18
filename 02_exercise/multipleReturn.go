package main

import "fmt"

func half(val int) (int, bool) {
	return val / 2, val%2 == 0
}

func main() {
	h, even := half(1)
	fmt.Println(h, even)
}
