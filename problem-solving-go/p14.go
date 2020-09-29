/*
	Problem :  Find out duplicate number between 1 to N numbers.
*/

package main

import "fmt"

//Return sum of given array
func sumArr(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func findDuplicate(arr []int) {
	n := len(arr)
	sum := sumArr(arr)
	total := ((n - 1) * n) / 2
	dup := sum - int(total)
	fmt.Println(dup)

}

func main() {
	arr := []int{1, 2, 2, 4, 5}
	findDuplicate(arr)

}
