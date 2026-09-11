package main

import (
	"bufio"
	"fmt"
	"os"
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

	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {

	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (A - add item, S - save bill, T - add tip) :", reader)
	fmt.Println(opt)
}

func main() {
	//User Input
	myBill := createBill()

	promptOptions(myBill)
}
