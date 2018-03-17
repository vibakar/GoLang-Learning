package main

import "fmt"

func main() {
	var lgNumber int
	var smNumber int
	fmt.Println("Enter the large number")
	fmt.Scan(&lgNumber)
	fmt.Println("Enter the small number")
	fmt.Scan(&smNumber)
	fmt.Println("The remainder of 2 number is ", lgNumber%smNumber)
}
