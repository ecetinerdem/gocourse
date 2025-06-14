package errorhandling

import (
	"fmt"
)

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d %s", b.parameter, b.message)
}

func Guess2(guessedNum int) (string, error) {
	if guessedNum < 1 || guessedNum > 100 {
		return "", &BorderException{guessedNum, "Out of border"}
	}
	return "You won", nil
}
