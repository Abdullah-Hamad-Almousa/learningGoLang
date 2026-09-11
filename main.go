package main

import "fmt"

func updateNames(x *string) {
	*x = "Ali"
}

func main() {
	//only about Pointers

	name := "Ahemd"

	m := &name
	//fmt.Println("Memeory address:", m)
	//fmt.Println("The value of the memory address is:", *m)
	//What we just do is make m as a pointer to use it with name to change it from type A to be a type B, to be a pointer
	fmt.Println(name)
	updateNames(m)
	fmt.Println(name)

}
