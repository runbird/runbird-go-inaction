package main

import "fmt"

type notifier interface {
	notify()
}

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

// ! 内部类的实现方法被提升到了外部类型
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{name: "scy",
			email: "scy@email.com",
		},
		level: "system",
	}
	//直接访问内部类型的方法
	//Sending user email to scy<scy@email.com>
	ad.user.notify()
	//内部类型方法提升到外部类型访问
	//Sending user email to scy<scy@email.com>
	ad.notify()

	//Sending user email to scy<scy@email.com>
	sendNotification(&ad)
}
