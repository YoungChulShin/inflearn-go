package main

import "fmt"

func if1() {
	// if문은 무조건 boolean으로 검사
	// 소괄호를 사용하지 않는다

	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}
}
