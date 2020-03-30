package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	if a == 10 {
		fmt.Printf("value of a is 10\n")
	} else if a == 20 {
		fmt.Printf("value of a is 20\n")
	} else if a == 30 {
		fmt.Printf("value of a is 30\n")
	} else {
		fmt.Printf("none of the values is matching\n")
	}
	if a == 100 {
		if b == 200 {
			fmt.Printf("value of a is 100 and b is 200\n")
		}
	}

	fmt.Printf("exact value of a is : %d\n", a)
	fmt.Printf("exact value of b is : %d\n", b)

	var grade string
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of x :%T\n", i)
	case int:
		fmt.Printf("x is int\n")
	case float64:
		fmt.Printf("x is float64\n")
	case func(int) float64:
		fmt.Printf("x is func(int)\n")
	case bool, string:
		fmt.Printf("x is bool or string\n")
	default:
		fmt.Printf("don't know the type\n")
	}

	var c1, c2, c3 chan int
	var i1, i2 int

	select {
	case i1 = <-c1:
		fmt.Printf("received %d from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d to c2\n", i2)
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received %d from c3\n", i3)
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}
