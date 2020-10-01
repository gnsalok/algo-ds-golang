package main

import "fmt"

func main() {
	slice := []string{"Alok", "Atul", "Anurag"}

	// Slicing using index
	fmt.Println(slice[0:2])

	//Appending
	slice = append(slice, "Dimple")
	fmt.Println(slice)

}
