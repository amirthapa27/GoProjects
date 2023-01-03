package main

import (
	"fmt"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	// panic("Please implement DescribeNumber")
	return fmt.Sprintf("This is the number %0.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	// panic("Please implement DescribeNumberBox")
	return fmt.Sprintf("This is a box containing the number %.d.0", nb.Number())
}

func main() {
	fmt.Println(DescribeNumber(-23))
	fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
}
