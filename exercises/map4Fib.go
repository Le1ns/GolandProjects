package main

import (
	"fmt"
)

//все ок
func main() {
	newFib := memoize(fib)
	for i := 1; i < 11; i++ {
		fmt.Println(newFib(i))
	}
}

func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if cache[n] == 0 {
			cache[n] = f(n)
			fmt.Println("----", cache[n])
		}
		return cache[n]
	}
}

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
