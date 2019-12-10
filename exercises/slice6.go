package main

import (
	"fmt"
	"os"
)

func deleteNumber(x int) ([]int, int) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}

	k := a[x]
	// Remove the element at index i from a.
	copy(a[x:], a[x+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = 0      // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a, k
}
func main() {
	//main
	var x int
	fmt.Print("vvdedite chislo ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Println(deleteNumber(x))
}
