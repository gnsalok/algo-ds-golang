/*
Problem : Find all pairs (a, b) in an array such that a+b=c+d]

*/

package main

import "fmt"

func findPairs(arr []int) {
	if arr == nil {
		fmt.Printf("Empty!")
	}

	n := len(arr)

	type values struct {
		v1, v2 int
	}

	m := make(map[int]values)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := arr[i] + arr[j]
			if val, found := m[sum]; found {
				fmt.Printf("%v && {%v %v}\n", val, arr[i], arr[j])
			} else {
				m[sum] = values{arr[i], arr[j]}
			}
		}
	}

}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	findPairs(arr)

}
