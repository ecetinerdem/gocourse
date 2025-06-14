package stringfunctions

import (
	"fmt"
	"strings"
)

func Demo2() {
	name := "Engin"

	fmt.Println(strings.HasPrefix(name, "Eng"))
	fmt.Println(strings.HasSuffix(name, "in"))

	letters := []string{"e", "n", "g", "i", "n"}

	fmt.Println(strings.Join(letters, ""))
	fmt.Println(strings.Replace(name, "E", "e", 1))

	fmt.Println(strings.Split(name, ""))

	fmt.Println(strings.Repeat(name, 5))
}
