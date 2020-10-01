package main

import "fmt"

//furling of array
func Sum(arr ...int) int {
	sum := 0

	for _, value := range arr {
		sum = sum + value
	}
	return sum

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//unfurling of array
	total := Sum(arr...)
	fmt.Println(total)

}
