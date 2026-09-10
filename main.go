package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {

	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], ""
	}

}

func main() {

	//Multiple Return

	vr, tr := getInitials("Ahmed lockhart")

	fmt.Println(vr, tr)

	vr2, tr3 := getInitials("Cloud fie")

	fmt.Println(vr2, tr3)

	vr3, tr3 := getInitials("Sun")

	fmt.Println(vr3, tr3)

}
