package errorhandling

import (
	"errors"
	"fmt"
)

func Guess(guessedNum int) (string, error) {
	if guessedNum < 1 || guessedNum > 100 {
		return "", errors.New("1-100 arasında bir sayı dönünüz")

	}
	return "Bildiniz", nil
}

func Demo2() {
	message, _ := Guess(10)

	fmt.Println(message)
	fmt.Println(Guess(2))
	fmt.Println(Guess(101))
	fmt.Println(Guess((-10)))

}
