package loops

import "fmt"

func DemoTwo() {

	var givenNumber int = 0

	fmt.Println("Please enter a number")
	fmt.Scanln(&givenNumber)

	if givenNumber <= 1 {
		fmt.Printf("%v, is not a prime number", givenNumber)
	}

	var isPrime bool = true

	for i := 2; i*i <= givenNumber; i++ {
		if givenNumber%i == 0 {
			isPrime = false
			break
		}
	}

	if !isPrime {
		fmt.Printf("%v, is not a prime number", givenNumber)
	} else {
		fmt.Printf("%v, is not a prime number", givenNumber)
	}

}
