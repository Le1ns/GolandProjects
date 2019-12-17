package main

import (
	"fmt"
	"strings"
	"unicode"
)
/// Супер!
func count(str string) map[string]int {

	slovar := make(map[string]int)
	converted := strings.FieldsFunc(str, func(с rune) bool {
		return !unicode.IsLetter(с)
	})
	println(converted)
	for _, w := range converted {

		slovar[w]++
	}
	return slovar
}

func main() {
	words := "aaaaaaaa bbbbbbbbbbb ccccccccccc dddddddddddd eee ffff eee"

	fmt.Println(count(words))

}
