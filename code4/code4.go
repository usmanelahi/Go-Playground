package main

import "fmt"

func stringLen(s1 string, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)

	return l1 + l2
}

func main() {
	s1 := "cristiano"
	s2 := "ronaldo"

	res := stringLen(s1, s2)

	fmt.Printf("The characters of string are %d \n ", res)
}
