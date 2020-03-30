package main

import "fmt"

func main() {
	var a int
	var b int = 15
	numbers := [6]int{1, 2, 3, 4, 5}
	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}
	for a < b {
		a++
		fmt.Printf("value of a: %d\n", a)
	}
	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x, i)
	}

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}

	var c int = 10
	for c < 20 {
		if c == 15 {
			c++
			continue
		}
		fmt.Printf("value of c: %d\n", c)
		c++
	}

	var d int = 10
LOOP:
	for d < 20 {
		if d == 15 {
			d++
			goto LOOP
		}
		fmt.Printf("value of d: %d\n", d)
		d++
	}
}
