package main

import "fmt"

func main() {
	fmt.Println("sample app")
}

func SampleMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
