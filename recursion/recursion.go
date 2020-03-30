package main

import "fmt"

func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	i := 15
	fmt.Println("factorial of", i, "is", factorial(i))

	fmt.Printf("fibonacci: ")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}
