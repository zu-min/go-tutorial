package main

import (
	"fmt"
	"math"
)

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func swap(x, y string) (string, string) {
	return y, x
}

func swap_address(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	x, y, radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("0x%x", int(h))
}

func main() {
	var a int = 100
	var b int = 200

	ret := max(a, b)
	fmt.Printf("max value is : %d\n", ret)

	x, y := swap("Mahesh", "Kumar")
	fmt.Println(x, y)

	fmt.Println(a, b)
	swap_address(&a, &b)
	fmt.Println(a, b)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	nextNum := getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	nextNum = getSequence()
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("circle area: %f\n", circle.area())

	var h Hex = 100
	fmt.Printf("hex: %s\n", h.String())
	fmt.Println(Hex(100))
}
