package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	c, d := swap("bonjour", "goodbye")
	fmt.Println(a, b, c, d)
}
