package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("My favorite number is", rand.Intn(10))
	}
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}
