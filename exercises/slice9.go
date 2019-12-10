package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	k := a[0]
	for i := 0; i < len(a)-1; i++ {
		temp := a[i+1]
		a[i] = temp
	}
	a[len(a)-1] = k
	fmt.Println(a)

}
