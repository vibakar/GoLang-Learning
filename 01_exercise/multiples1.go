package main

import "fmt"

var sum int;

func main() {
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	fmt.Println("Sum is ",sum)
}

