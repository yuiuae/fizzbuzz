package main

import "fmt"

func fb(i int) {
	switch {
	case i%15 == 0:
		fmt.Println("fizzbuzz")
	case i%3 == 0:
		fmt.Println("fizz")
	case i%5 == 0:
		fmt.Println("buzz")
	default:
		fmt.Println(i)
	}
}
func main() {
	for i := 1; i < 100; i++ {
		fb(i)
	}
}
