package main

import "fmt"

func if2() {
	var a int = 50
	b := 70

	if a >= 65 {
		fmt.Println("65이상")
	} else {
		fmt.Println("65이하")
	}

	if b >= 70 {
		fmt.Println("70이상")
	} else {
		fmt.Println("70미만")
	}
}
