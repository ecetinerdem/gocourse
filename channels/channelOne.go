package channels

func EvenNums(evenChannel chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum += i
	}
	evenChannel <- sum
}

func OddNums(oddsChannel chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum += i
	}
	oddsChannel <- sum
}
