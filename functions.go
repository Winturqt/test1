package main

import "fmt"

func add(x int, y int) int {
	return (x + y) * 5
}

func main() {
	fmt.Println(add(25, 23))
}
