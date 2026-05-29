//go:build !skip
// +build !skip

package main

import (
	"fmt"
)

// This is like dictionary

func learn_map() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 2

	for k, v := range m {
		fmt.Println("The key value", k, v)
	}
}

type User struct {
	Name  string
	Email string
	Age   int
}

func update_email(u *User, new_email string) {
	u.Email = new_email
}

type Learner interface {
	learn(subject string)
}

func (u *User) learn(subject string) {
	fmt.Println("%s is learning %s \n", u.Name, subject)
}

func main() {

	learn_map()

	u := User{
		Name:  "Raj",
		Email: "raj@example.com",
		Age:   30,
	}
	fmt.Println(u.Name, u.Email, u.Age)
	update_email(&u, "raj@abc.com")
	fmt.Println(u.Name, u.Email, u.Age)
	u.learn("GO")
}
