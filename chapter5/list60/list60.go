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

//定义外部类的实现方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name, a.email)
}

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

	//ad类实现全部重写
	//内部类型方法提升到外部类型访问
	//Sending admin email to scy<scy@email.com>
	ad.notify()

	//Sending admin email to scy<scy@email.com>
	sendNotification(&ad)
}
