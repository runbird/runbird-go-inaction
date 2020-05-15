package main

import (
	"fmt"
	"github.com/goinaction/code/chapter5/listing74/entities"
)

func main() {
	//公开类型的Admin
	//嵌入未公开的user
	//未公开的user属性全公开
	//由于内部类user是未公开的，所以不能直接通过结构字面量的方式来初始化该内部类型
	ad := entities.Admin{
		Rights: 10,
	}
	//有点类似于接口，内部类型的标识符提升到了外部类型，公开的字段也可以通过外部类型的字段值来访问
	//不能直接访问内部类型，user.Name  user.Email，需要外部来访问
	ad.Name = "Bill"
	ad.Email = "bill@email.com"

	fmt.Printf("User: %v\n", ad)
}
