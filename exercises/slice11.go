package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := a[len(a)-1]
	for i := len(a) - 1; i != 0; i-- {
		temp := a[i-1]
		a[i] = temp
	}
	a[0] = k
	fmt.Println(a)

}
