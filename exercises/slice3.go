package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{5}
	b = append(b, a...)
	fmt.Println(b)

}
