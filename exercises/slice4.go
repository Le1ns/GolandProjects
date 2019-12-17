package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	i := len(a) - 1
	k := a[i]
	a = a[:len(a)-1]
	//не очень понятно, зачем нужен был этот кусок
	fmt.Println(a, k) // [A B D E]

}
