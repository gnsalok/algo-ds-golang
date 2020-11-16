// Reverse the array

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5}
	reverseArray(arr)

}

func reverseArray(a []int) {
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Print(a[i])
		if i > 0 {
			fmt.Print(", ")
		}
	}

}
