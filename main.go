package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//Standard Library

	greeting := "Hello my nerds!"
	fmt.Println(strings.Contains(greeting, "nerds!"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.ToLower(greeting))

	fmt.Println(strings.Index(greeting, "my"))

	fmt.Println(strings.Count(greeting, "my"))
	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("original string value =", greeting)

	ages := []int{10, 20, 30, 40, 50, 25, 60, 90, 52, 52}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 52)

	fmt.Println(index)

	index = sort.SearchInts(ages, 99)

	fmt.Println(index)

	names := []string{"Ahmed", "Abdullah", "Hamad", "Ali", "Mohammed"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Ali"))

}
