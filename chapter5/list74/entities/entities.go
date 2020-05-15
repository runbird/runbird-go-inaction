package entities

//未公开类，属性全公开
type user struct {
	Name  string
	Email string
}

type Admin struct {
	user   //嵌入
	Rights int
}
