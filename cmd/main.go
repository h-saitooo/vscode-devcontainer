package main

import (
	"fmt"
	"go-dev-container/pkg/compute"
)

func main() {
	fmt.Println("Hello from DevContainer")

	var a int = 4
	var b int = 9
	result := compute.SumNumber(a, b)

	fmt.Printf("Sum number %v, %v, result %v \n", a, b, result)
}
