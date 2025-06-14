package stringfunctions

import (
	"fmt"
	"strings"
)

func Demo1() {
	name := "Engin"
	//case sensitive
	fmt.Println(strings.Count(name, "n"))
	fmt.Println(strings.Contains(name, "n"))
	fmt.Println(strings.Index(name, "n"))
}
