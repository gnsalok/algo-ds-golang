package main

import "fmt"

func main() {

	i := 0
	a := []int{1, 2, 3}
	var mergedArr []int

	for c := 0; c < len(a); c++ {
		mergedArr = append(mergedArr, a[i:2]...)
		break
	}
	fmt.Println(mergedArr)

}
