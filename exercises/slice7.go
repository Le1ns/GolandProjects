package main

import "fmt"

//Объединить два слайса и вернуть новый со всеми элементами первого и второго

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	c := append(a, b...)
	fmt.Println(c)

}
