package main

import "fmt"

func main() {
	//Receiver Functions with Pointers

	myBill := newBill("Abdullah's bill")

	myBill.updateTip(10)
	myBill.addItem("Coffee", 12)
	myBill.addItem("Milk", 2)
	myBill.addItem("Cube-Cake", 15)

	fmt.Println(myBill.formatBill())
}
