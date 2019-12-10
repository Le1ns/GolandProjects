package main

import (
	"fmt"

)

func main() {
	p := []int{1,2,3,4,5,6}
	for v := range p {
		p[v]=p[v]+1
	}
	fmt.Println(p)
}
