package main

import (
	"fmt"
	_ "os"
)

//12e задание?

func main() {
	//var k int
	//fmt.Println("vvedite razmer slice")
	//	fmt.Scan(os.Stdin,&k)
	kek := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var kek string
	//fmt.Fscan(os.Stdin,&kek)
	vupsen := make([]int, len(kek), (cap(kek)+1)*2)
	copy(vupsen, kek)
	//kek=vupsen
	fmt.Println(vupsen)

}
