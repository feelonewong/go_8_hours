package main

import "fmt"

type User struct {
	UserName    string
	Age         int
	Address     string
	PhoneNumber int
}

func main() {
	user := User{
		UserName:    "huangfeilong",
		Age:         20,
		Address:     "安徽省淮北市濉溪县",
		PhoneNumber: 173561,
	}
	fmt.Printf(" user type: %T, user is: %v\n", user, user)
}
