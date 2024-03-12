package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 5
	y := 5
	var res = add(x, y)
	fmt.Printf("the result is %d\n", res)
}
