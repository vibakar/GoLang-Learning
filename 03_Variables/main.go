package main

import "fmt"

var name string = "viba"
var c int
var d int = 10
var memory *int = &d
var m = &name

const (
	num1 int    = 1
	num2 string = "one"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) func() int {
	return func() int {
		return a * b
	}
}

func main() {
	a := 8
	fmt.Println(name, a, c, d)
	*m = "virat"
	fmt.Println(name)
	fmt.Println(add(2, 3))
	fmt.Printf("%p \n", &a)
	*memory = 23
	fmt.Println(d)
	ans := mul(2, 3)
	fmt.Println(ans())
	fmt.Println(num1, num2)
}
