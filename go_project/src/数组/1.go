package main

import "fmt"

func main() {
	var numbers = [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[1])
	fmt.Println(numbers[len(numbers)-1])
}
