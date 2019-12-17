package main

import (
	_ "crypto/aes"
	"fmt"
	//	"math/rand"
	//"time"
)

/*func makefull(arr []int)[]int {
	rand.Seed(time.Now().Unix()) //eto obnovlenie vremeni????no teper rabotaet random
	test := make([]int, len(arr), cap(arr))
	lenght := len(test)
	for i := 0; i < len(test); i++ {
		test[i] = rand.Intn(lenght)
	}
	return test
}*/

func Onlyone(arr1, arr2 []int) map[int]int {
	massiv1 := make(map[int]int)
	massiv2 := make(map[int]int)
	//res:=make(map[int]int)
	for _, i := range arr1 {
		massiv1[i]++
	}
	for _, i := range arr2 {
		massiv2[i]++
	}

	for k, i := range massiv1 {
		for j, n := range massiv2 {
			if k == j {
				fmt.Println(k, i, "=", j, n)

			}

		}
	} //selano izi(net)
	return massiv1
}

func main() {
	mas1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mas2 := []int{7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	//	fmas1 := makefull(mas1)
	//	rand.Seed(time.Now().Unix())
	//fmas2:=makefull(mas2)
	Onlyone(mas1, mas2)
	//	fmt.Println(fmas1)
	//fmt.Println(fmas2)

}
