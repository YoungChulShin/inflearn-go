package main

import "fmt"

func main() {

	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	fmt.Println("shortvar1: ", shortVar1, " shortvar2: ", shortVar2, " shortvar3: ", shortVar3)

	if i := 10; i < 9 {
		fmt.Println("성공: ", i)
	} else {
		fmt.Println("실패: ", i)
	}
}
