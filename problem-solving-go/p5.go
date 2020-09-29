/*
Problem : Find the missing noumbers from 1 to 20.
If given array is sorted.
*/
package main

import "fmt"

func main() {

	arr := []int{1, 3, 4, 5, 7, 8, 10} // 2, 6

	n := len(arr)

	var missing []int

	for i := 1; i < n; i++ {
		if (arr[i] - arr[i-1]) != 1 {
			val := arr[i] - arr[0]
			missing = append(missing, val)
		}
	}
	fmt.Println(missing)

}
