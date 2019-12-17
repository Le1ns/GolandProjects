package main

import (
	_ "crypto/aes"
	"fmt"
	"math/rand"
	"time"
)

//зачем передавать весь массив, если используешь только длину его. Можно было сразу длину передать
func makefull(arr []int) []int {
	// Это установки для генератора случайных чисел
	rand.Seed(time.Now().Unix()) //eto obnovlenie vremeni????no teper rabotaet random
	test := make([]int, len(arr), cap(arr))
	// эквивалентно
	//test := make([]int, len(arr))
	lenght := len(test)
	for i := 0; i < len(test); i++ {
		test[i] = rand.Intn(lenght)
	}
	return test
}

// имена функций в go должны следовать CamelCase нотации. Соответственно OnlyOne.
func Onlyone(arr []int) map[int]int {
	massiv := make(map[int]int)
	for _, i := range arr {
		massiv[i]++
	}

	for k, i := range massiv {
		//Хоть по разу. i>=1
		//Обычно i - это индекс v - значение. Когда они поменяны местами - сибивает.
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
