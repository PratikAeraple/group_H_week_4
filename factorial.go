package main

// Function to calculate the factorial of a number
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}
