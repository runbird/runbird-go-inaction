package main

import "fmt"

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

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@emali.com"}
	sendNotification(&u)
	// sendNotification(u)
	// 不能将u 作为sendNotification的参数类型notifier
	//user 类型并没有实现notifier （notify方法使用指针接收者声明）
	fmt.Println(u)
}

// 方法集
// Values  Methods Receivers
// T        (t T)
// *T       (t T) and (t *T)

// Methods Receivers        Values
// (t T)                    T and *T
// (t *T)                   *T
