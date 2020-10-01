package main

import "fmt"

type User struct {
	Id    int
	Name  string
	Email string
}

func (user *User) Get() {
	// Get all the data of User
	user.Id = 101
	user.Name = "Alok Tripathi"
	user.Email = "gnsalok@gmail.com"
	fmt.Println(user.Id, user.Name, user.Email)
}

func main() {

	u := User{Id: 110}
	u.Get()

	// Will get everything here, when  do pass by ref
	fmt.Println(u.Id, u.Name, u.Email)

}
