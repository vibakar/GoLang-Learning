package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("What is your age?")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are eligible for vote")
	} else {
		fmt.Println("You are a kiddo.You are not eligible for vote")
	}

}
