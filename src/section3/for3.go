package main

import "fmt"

func for3() {

	// case 1
Loop1: // loop 아래에는
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1 // 가장 바깥으로
			}
			fmt.Println(i, j)
		}
	}

	fmt.Println("====================")

	// case 2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("====================")

Loop2:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				continue Loop2
			}
			fmt.Println(i, j)
		}
	}
}
