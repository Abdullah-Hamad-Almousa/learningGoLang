package main

import "fmt"

func updateName(x string) string { //It takes a copy of original variable that why it gonna print a Ahmed
	x = "Fahad"

	return x
}

func updateMenu(y map[string]int8) { //It gonna pass the original variable that why it gonna print the new value
	y["Coffee"] = 10
}

func main() {
	//Passing
	//Group type A -> strings, int, float, bools, arrays and structs

	//name := "Ahemd"
	//
	//name = updateName(name)
	//
	//fmt.Println(name)

	//Group type B -> slice, maps and functions

	menu := map[string]int8{
		"Pie":       24,
		"Salad":     18,
		"Ice cream": 2,
	}

	updateMenu(menu)

	for k, v := range menu {
		fmt.Println(k, v)
	}

	//Type A can't be changed from outside it area unless used return
	//Type B can be changed from outside it area
	//Type A can be called non-pointer values and B can be called pointer values

}
