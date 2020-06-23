package main

import (
	"fmt"
)

func add(a, b int) int { return a + b }

func negative(a, b int) int { return a - b }

func kun(a, b int) int { return a * b }

func haft(a, b int) int { return a / b }

func main() {
	a := 4
	b := 2
	c := add(a, b)
	e := negative(a, b)
	f := kun(a, b)
	g := haft(a, b)
	fmt.Println("add")
	fmt.Printf("result = %d", c)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("negative")
	fmt.Printf("result = %d", e)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("kun")
	fmt.Printf("result = %d", f)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("haft")
	fmt.Printf("result = %d", g)
}
