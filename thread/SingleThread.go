package main

import (
	"fmt"
	"math/rand"
)

func rand_generate_1() int {
	return rand.Int()
}

func main() {
	fmt.Println(rand_generate_1())
	fmt.Println(rand.Float32())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Int63n(100000))

	fmt.Println(rand.ExpFloat64())
}
