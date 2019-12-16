package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(slice)-1; i+=2 {
		slice[i], slice[i+1] = slice[i+1], slice[i]
	}
	fmt.Println(slice)

}
