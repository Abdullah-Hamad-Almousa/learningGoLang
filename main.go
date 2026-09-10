package main

import "fmt"

func main() {

	var age int8 = 27
	fName := "Abdullah"
	lName := "Almousa"
	var experince byte = 3

	fmt.Println("Hello", fName, lName, "with experince of:", experince, "years and age:", age)
	fmt.Printf("Hello %v %v with experince of: %v years and age: %v \n", fName, lName, experince, age)
	fmt.Printf("Hello %q %q with experince of: %q years and age: %q \n", fName, lName, experince, age)
	fmt.Printf("Age is of type %T\n", age)
	fmt.Printf("you scored %f points! \n", 99.4)
	fmt.Printf("you scored %.2f points! \n", 99.4)

	// Sprintf

	var strOne = fmt.Sprint("Hello ", fName, " ", lName, " with experince of: ", experince, " years and age: ", age)
	var strTwo = fmt.Sprintf("Hello", fName, lName, "with experince of:", experince, "years and age:", age)
	fmt.Println("String have been saved:", strOne, "\n", strTwo)

}
