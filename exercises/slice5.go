package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := 0
	k := a[i]
	a = a[1:]
	fmt.Println(a, k) // [A B D E]
}
