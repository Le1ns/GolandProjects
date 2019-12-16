package main

import (
	"fmt"
	"sort"
)

//Объединить два слайса и вернуть новый со всеми элементами первого и второго

//не рабочая.
//Вернула [2 3 4 6 7 8 9 10]
func _main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 12, 13, 5, 15, 16, 6, 18, 7, 20}

	for i := 0; i < len(a); i++ {

		if a[i] == b[i] {
			copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
			a[len(a)-1] = 0      // Erase last element (write zero value).
			a = a[:len(a)-1]
		}

	}

	fmt.Println(a)
}


func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{1, 12, 13, 5, 15, 16, 6, 18, 7, 20}
	res:=make([]int, 0, len(a))
	sort.Ints(b)
	for i:=range a {
		if index:=sort.SearchInts(b, a[i]); index < len(b) && b[index]!=a[i] {
			res=append(res, a[i])
		}
	}
	a=res
	fmt.Println(a)
}
