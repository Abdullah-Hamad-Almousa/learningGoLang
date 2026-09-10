package main

import "fmt"

func main() {

	//Arrays and Slices

	//Arrays
	var ages [3]int = [3]int{25, 27, 29}

	names := [4]string{"Ali,", "Hamad", "Fahad", "Mohammad"}
	names[0] = "Ahmed"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//Slices
	var scores = []int{93, 95, 99}
	scores[2] = 97
	scores = append(scores, 99) //append is overwrite on the Slices and adding a new value with it

	fmt.Println(scores, len(scores))

	//Slices ranges

	rangeOne := names[1:3]  //Print between index 1 and 3. 1 is gonna print but it not gonna print 3
	rangeTwo := names[2:]   //Print everything on index 1 to the last
	rangeThree := names[:2] //Print everything from the start until index 2
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Abdullah") //Is not gonna add it to array names it only gonna add it to rangeOne
	fmt.Println(rangeOne)
}
