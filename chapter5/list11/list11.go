package main

import "fmt"

type user struct {
	name  string
	email string
}

// notify使用值接收者实现
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

//changeEmail 使用指针接收者实现
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@mail.com"}
	bill.notify()

	lisa := &user{"Lisa", "lisa@email.com"}
	//(*lisa.notify())
	lisa.notify()

	bill.changeEmail("bill@google.com")
	bill.notify()

	lisa.changeEmail("lisa@google.com")
	lisa.notify()
}
