package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) { //We got bufio.Reader by hover on reader in createBill()
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) //Call for the input reader

	//fmt.Println("Create a new bill name:")
	//name, _ := reader.ReadString('\n') //The trigger for reading next line of code when the user click enter by define \n
	//name = strings.TrimSpace(name)     //Delete space that are not between the string

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (A - add item, S - save bill, T - add tip) : ", reader)

	switch opt {
	case "A":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)
	case "S":
		fmt.Println("You chose to save the bill", b)
	case "T":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added -", tip)
		promptOptions(b)
	default:
		fmt.Println("That's not a valid option")
		promptOptions(b)
	}
}

func main() {
	//Parsing Floats
	myBill := createBill()

	promptOptions(myBill)
}
