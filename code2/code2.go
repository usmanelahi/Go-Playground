// Write a code to print Even numbers in given array
package main

import "fmt"

func main() {
	temp := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(temp); i++ {
		if temp[i]%2 == 0 {
			fmt.Printf("Even Number %d\n", temp[i])
		}
	}

}
