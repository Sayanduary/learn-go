package main

import (
	"fmt"
	"my_project/auth"
	"my_project/user"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("sayan", "1234")
	fmt.Println(auth.GetSession())
	user :=
		user.User{
			Email: "sayanduary@gmail.com",
			Name:  "Sayan Duary",
		}
	fmt.Println(user.Email)
	color.Green(user.Name)

}
