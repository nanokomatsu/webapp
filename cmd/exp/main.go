package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Bobby Hill",
		Bio:  `<script>alert("Hehe, you have been h4x0r3d!");</script>`,
		Age:  13,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
