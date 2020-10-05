package main

import "fmt"

func mergeSortedArray(arr1 []int, arr2 []int) []int {

	var mergedArr []int
	lengthArr := len(arr1) + len(arr2)
	fmt.Println(lengthArr)

	i := 0
	j := 0
	//Check input
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}
	for c := 0; c < lengthArr; c++ {
		if i >= len(arr1) {
			mergedArr = append(mergedArr, arr2[j:len(arr2)]...)
			break
		} else if j >= len(arr2) {
			mergedArr = append(mergedArr, arr1[i:len(arr1)]...)
			break
		} else if arr1[i] >= arr2[j] {
			mergedArr = append(mergedArr, arr2[j])
			j++
		} else {
			mergedArr = append(mergedArr, arr1[i])
			i++
		}
	}

	return mergedArr

}

func main() {

	arr1 := []int{0, 3, 31}
	arr2 := []int{4, 6, 30, 45}

	m := mergeSortedArray(arr1, arr2)
	fmt.Println(m)
	//Exp output : 0,3,4,6,30,31
}
