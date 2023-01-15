package main

import "fmt"

func switch2() {
	// rand.Seed(time.Now().UnixNano())
	// switch i := rand.Intn(100); {
	// case i >= 50 && i < 100:
	// 	fmt.Println("i -> ")
	// }

	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("짝수")
	case 1, 3, 5:
		fmt.Println("홀수")
	default:
		fmt.Println("정해진 값이 아닙니다")
	}

	// go는 자동 break;
	// fallthrough
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
	case "go":
		fmt.Println("go!")
		fallthrough // 다음 케이스로 넘어간다. 조건이 맞지 않아도 실행
	case "python":
		fmt.Println("python!")
	case "kotlin":
		fmt.Println("kotlin!")
	case "csharp":
		fmt.Println("csharp!")
	}
}
