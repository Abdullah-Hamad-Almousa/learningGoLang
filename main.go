package main

import "fmt"

func main() {
	//Structs and custom

	myBill := newBill("Abdullah's bill")

	//myBill.formatBill() // formatBill it is not a function with by self anymore

	fmt.Println(myBill.formatBill())
}
