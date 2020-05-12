package main

import "fmt"

//多态 A、B 实现了 接口 notifier
//多态函数 sendNotification( n notifier) ;

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@emali.com"}
	sendNotification(&u)

	a := admin{"Lisa", "lisa@email.com"}
	sendNotification(&a)
}
