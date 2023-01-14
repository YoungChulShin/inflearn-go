package main

import (
	"fmt"
)

func main() {
	const (
		E = 2
		A = 1
		_
		C = iota
	)

	fmt.Println(E, A, C)
}
