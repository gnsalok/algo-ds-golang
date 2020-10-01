/*
Problem : Find the missing noumbers from 1 to 20.
If given array is sorted.
*/
package main

import "fmt"

func main() {

	arr := []int{4, 5, 7, 8, 10, 20} // 2, 6
	one := 1

	n := len(arr)

	var missing []int

	for i := 1; i < n; i++ {
		if (arr[i] - arr[i-1]) != 1 {
			val := arr[i] - one
			missing = append(missing, val)
		}
	}
	fmt.Println(missing)

}
