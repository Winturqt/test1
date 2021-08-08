package main

import "fmt"

func split(sum int) (w, x, y, z int) {
	w = sum * 5 / 10
	x = ((sum - w) / 2) + 1
	y = x/2 - 1
	z = y/2 - 1
	return
}

func main() {
	fmt.Println(split(100))
}
