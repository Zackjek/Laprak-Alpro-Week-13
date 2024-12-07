package main

import "fmt"

func main() {
	var number int
	var cotinueLoop bool
	for cotinueLoop = true; cotinueLoop; {
		fmt.Scan(&number)
		cotinueLoop = number <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", number)
}