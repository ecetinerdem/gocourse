package main

import (
	"fmt"
	"golesson/channels"
)

func main() {
	//variables.Variables()

	//conditionals.Conditionals()
	//conditionals.ConditionalsTwo()
	//conditionals.Demo()

	//loops.Loops()
	//loops.Demo()
	//loops.DemoTwo()
	//loops.DemoThree()

	//arrays.Arrays()
	//arrays.ArraysTwo()
	//arrays.ArraysThree()
	//arrays.ArraysFour()

	//sliceOne.SlicesOne()
	//sliceOne.SlicesTwo()

	// functions.SayHi("Cetin")
	// functions.Sum(2, 6)
	//sum1, subtract1, multiply1, divide1 := functions.FourOperations(6, 2) underscore for not having any of the values
	//var result = functions.VariadicSum(1, 4, 5, 3, 10)

	//forrange.RangeOne()
	//forrange.RangeTwo()

	//number sent to func is different address than number created!!
	// number := 20
	// pointers.PointerOne(&number)
	// fmt.Printf("Num in main: %v\n", number)

	// numbers := []int{1, 2, 3}

	// pointers.PointerTwo(numbers)

	// fmt.Printf("Num in main: %v\n", numbers[0])

	//forstruct.StructOne()

	//forstruct.StructTwo()

	//go goroutines.EvenNums()
	//go goroutines.OddNums()
	//time.Sleep(5 * time.Second)

	evenSumChannel := make(chan int)
	oddsSumChannel := make(chan int)
	go channels.EvenNums(evenSumChannel)
	go channels.OddNums(oddsSumChannel)

	evenSum, oddSum := <-evenSumChannel, <-oddsSumChannel

	multiply := evenSum * oddSum

	fmt.Println("Çarpım: ", multiply)
}
