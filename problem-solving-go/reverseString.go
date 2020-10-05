package main

import "fmt"

func reverseString(str string) (result string) {

	for _, v := range str {
		//keep appending the char
		result = string(v) + result
		fmt.Println(result)
	}
	return

}

func main() {

	str := "Alok"
	fmt.Println(str)
	fmt.Println(reverseString(str))

}
