package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {

	fmt.Printf("Good morning %v \n", n)

}

func sayBye(n string) {

	fmt.Printf("Goodbye %v \n", n)

}

func cycleNames(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}

}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	//Functions

	sayGreeting("Abdullah")
	sayBye("Abdullah")

	cycleNames([]string{"Ahmed", "Ali", "Abdullah"}, sayGreeting)
	cycleNames([]string{"Ahmed", "Ali", "Abdullah"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Printf("Circle 1 is %.3f, and the circle 2 is %.3f", a1, a2)
	fmt.Printf("Circle 1 is %.3v, and the circle 2 is %.3v", a1, a2)

}
