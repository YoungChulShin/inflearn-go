package main

import "fmt"

func main() {
	const a, b int = 10, 99
	const c, d, e = true, 12.3, "test"
	const (
		x, y int16 = 50, 60
		i, k       = "Data", 1234
	)

	fmt.Println(a, b, c, d, e)
	fmt.Println(x, y, i, k)
}
