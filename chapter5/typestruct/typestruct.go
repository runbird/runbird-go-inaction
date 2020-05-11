package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

//var 声明变量并初始化
var bob user

//类似于 :=
func foo() {
	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Println(lisa)

	//user2{
	//	name:       "Lisa",
	//	email:      "lisa@email.com",
	//	ext:        123,
	//	privileged: true,
	//}
	//fmt.Println(user2)

	jim := user{"jim", "jim@qq.com", 1234, true}
	fmt.Println(jim)
}

type admin struct {
	person user
	level  string
}

func foo2() {
	adminis := admin{
		person: user{
			name:       "Oracle",
			email:      "oo@qq.com",
			ext:        0,
			privileged: false,
		},
		level: "admin",
	}
	fmt.Println(adminis)
}

//基于int64声明新类型
type Duration int64

func foo3() {
	//编译报错，不同类型
	//var dur Duration
	//dur = int64(1000)
}

func main() {

}
