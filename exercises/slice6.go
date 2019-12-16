package main

import (
	"fmt"
	"os"
)

func deleteNumber(x int) ([]int, int) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	if x>=len(a) {
		return a, 0
	}

	k := a[x]
	a = append(a[:x],a[x+1:]...)
	return a, k
}
func main() {
	//main
	var x int
	fmt.Print("vvdedite chislo ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Println(deleteNumber(x))
}
