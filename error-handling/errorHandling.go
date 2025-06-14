package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//Dosya bulunursa err == nil
	//f dosyamız

	if err != nil {
		fmt.Println("File cannot be found")
	} else {
		fmt.Println("File can be used with file functions", f.Name())
	}
}
