package main

import (
	_ "crypto/aes"
	"fmt"
	"math/rand"
	"time"
)

func makefull(arr []int) []int {
	rand.Seed(time.Now().Unix()) //eto obnovlenie vremeni????no teper rabotaet random
	test := make([]int, len(arr), cap(arr))
	lenght := len(test)
	for i := 0; i < len(test); i++ {
		test[i] = rand.Intn(lenght)
	}
	return test
}

func Onlyone(arr []int) map[int]int {
	massiv := make(map[int]int)
	for _, i := range arr {
		massiv[i]++
	}

	for k, i := range massiv {
		if i == 1 {
			fmt.Println(k, i)
		}
	} //selano izi(net)
	return massiv
}

func main() {
	mas := make([]int, 100, 100)
	c := makefull(mas)
	fmt.Println(Onlyone(c))

}
