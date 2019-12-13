package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceint := sort.IntSlice{4, 1, 3, 2}
	slicestring := []string{"Nikita", "Andrey", "Dima"}

	sort.Ints(sliceint)
	fmt.Println(sliceint)
	sort.Sort(sort.Reverse(sliceint))
	fmt.Println(sliceint)
	sort.Strings(slicestring)
	fmt.Println(slicestring)

}
