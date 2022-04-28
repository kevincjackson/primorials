package main

import (
	"fmt"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}

type primorial struct {
	id        int
	primes    []int
	product   int
	remaining []int
	size      int
}

func newPrimorial(x int) {
	id := x
	primes := primes[0:x]
	product := 1
	for i := 0; i < len(primes); i++ {
		product *= primes[i]
	}
	remaining := []int{}
outer:
	for i := 1; i <= product; i++ {
		for j := 0; j < id; j++ {
			if i%primes[j] == 0 {
				continue outer
			}
		}
		remaining = append(remaining, i)
	}
	size := len(remaining)
	percentagereduced := 1.0 - float32(size)/float32(product)
	// fmt.Println(id, primes, product, remaining, size, percentagereduced)
	fmt.Println(id, primes, size, product, percentagereduced)
	fmt.Println()
}

func main() {
	for i := 8; i <= 10; i++ {
		newPrimorial(i)
	}
}
