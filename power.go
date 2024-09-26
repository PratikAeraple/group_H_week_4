package main

// Function to calculate a number raised to the power of another number
func Power(base int, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
