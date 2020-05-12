package main

import "fmt"

//嵌入类型 复用
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

type admin struct {
	user
	level string
}

func main() {

}
