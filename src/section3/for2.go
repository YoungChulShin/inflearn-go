package main

import "fmt"

func for2() {

	// case 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1
	}
	fmt.Println(sum1)

	fmt.Println("====================")

	// case 2
	// go에서 후취연산은 반환 값 x
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
	}

	fmt.Println(sum2)

	fmt.Println("====================")
	// case 3 - while 문 형태
	sum3, i := 0, 0
	for {
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println(sum3)

	fmt.Println("====================")

	// case4 -
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}

}
