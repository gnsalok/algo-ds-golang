package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		n = n * fact(n-1)
	}
	return n
}

func main() {
	fmt.Println("Implement your solution")

	fmt.Println(fact(5))

}
