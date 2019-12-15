package main

import (
	_ "crypto/aes"
	"fmt"
	"math/rand"
	"sort"
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
	keys := []int{}
	for key := range massiv {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for i := 0; i < len(massiv); i++ {
		if keys[i] == 1 {
			fmt.Println(massiv[keys[i]])
		}
	} //ne dodelano
	return massiv
}

func main() {
	mas := make([]int, 100, 100)
	c := makefull(mas)
	fmt.Println(Onlyone(c))

}
