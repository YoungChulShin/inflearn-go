package main

import "fmt"

func for1() {

	// case 1
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("====================")

	// case2 - 무한루프
	// for {
	// 	fmt.Println("Ininite loop")
	// }

	fmt.Println("====================")

	// case 3 - ranage 사용
	// 첫번째가 index, 두번째가 값
	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc {
		fmt.Println(index, name)
	}
	for _, name := range loc {
		fmt.Println(name)
	}

	fmt.Println("====================")
}
