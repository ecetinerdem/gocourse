package functions

func FourOperations(num1 int, num2 int) (int, int, int, float32) {
	sum := num1 + num2

	subtract := num1 - num2

	multiply := num1 * num2

	divide := float32(num1 / num2)

	return sum, subtract, multiply, divide
}
