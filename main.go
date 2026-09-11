package main

import "fmt"

func main() {
	//Maps

	menu := map[string]float32{
		"pie":   24.3,
		"water": 2.1,
		"soup":  15.2,
		"salad": 18.2,
		"cream": 9.1,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu { // k will hold the string as index and the v will hold the price of the string
		fmt.Println(k, ":", v)
	}

	//ints as key

	phonebook := map[int]string{
		911: "Emergency",
		999: "Police",
		998: "Civil Defense",
		997: "Medical services",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[911])

	phonebook[911] = "Superman"
	fmt.Println(phonebook[911])

	for k, v := range phonebook {
		fmt.Println(k, ":", v)
	}

}
