package main

import "fmt"

func Factorial(n int) (result int) {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return n * Factorial(n-1)
}

func main() {
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
}
