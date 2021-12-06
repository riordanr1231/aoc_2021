package main

import (
	_ "embed"
	"fmt"
)

func main() {
	matrix := make([][]string, 5)
	for i := 0; i < 5; i++ {
		matrix[i] = make([]string, 5)
	}

	matrix[0][0] = "1"
	matrix[0][1] = "2"
	fmt.Println(matrix)
}
