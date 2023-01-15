package main

import (
	"fmt"
)

func switch1() {
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	// 이렇게 선언을 더 추천
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	// case 3
	switch c := "go"; c {
	case "go":
		fmt.Println("go")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// case 4
	switch c := "go"; c + "lang" {
	case "go":
		fmt.Println("go")
	case "java":
		fmt.Println("java")
	default:
		fmt.Println("일치하는 값 없음")
	}

	// case 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i가 j보다 작다")
	case i == j:
		fmt.Println("i가 j와 같다")
	case i > j:
		fmt.Println("i가 j보다 크다")
	}
}
