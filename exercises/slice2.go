package main

import (
	"fmt"
)

func main() {
	p := []int{1, 2, 3, 4, 5, 6}
	p = append(p, 5)
	fmt.Println(p)
}
