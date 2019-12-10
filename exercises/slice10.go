package main

import (
	"fmt"
	"os"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int
	fmt.Println("vvedite na skolko sdvinyt")
	fmt.Fscan(os.Stdin, &i)
	if i > 9 {
		fmt.Println("i slishkom bolshaya")
	}

	c := append(a[i:], a[:i]...)

	fmt.Println(c)

}
