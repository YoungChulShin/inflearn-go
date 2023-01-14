package main

import "fmt"

func main() {
	// 여러개를 묶어서 사용할 수 있다
	var (
		name      string = "machine"
		height    int32
		weight    float32
		isRunning bool
	)

	height = 250
	weight = 350.55
	isRunning = true

	fmt.Println("name : ", name, ", height : ", height, ", weight : ", weight, ", isRunning : ", isRunning)
}
