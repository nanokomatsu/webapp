package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta struct {
		Visits int
	}
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := User{
	// 	Name: "Rusty Shackleford",
	// }

	// user := struct {
	// 	Name string
	// 	Age  int
	// }{
	// 	Name: "John Redcorn",
	// 	Age:  40,
	// }

	user := User{
		Name: "Bobby Hill",
		Age:  13,
		Meta: UserMeta{
			Visits: 6,
		},
	}

	fmt.Println(user.Meta.Visits)

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
