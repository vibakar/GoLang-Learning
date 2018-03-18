package main

import "fmt"

func findGreatest(list ...int) int {
	var lgNum int
	for _, val := range list {
		if val > lgNum {
			lgNum = val
		}
	}
	return lgNum
}

func main() {
	fmt.Println(findGreatest(23, 5, 18, 29, 24, 27))
}
