package main

import (
	"fmt"
)

func main() {
	//声明一个包含5个元素的整型数组，初始值都为0
	var array1 [5]int
	fmt.Println(array1)

	// .... ,并赋值
	array2 := [5]int{10, 20, 30, -10, 50}

	//容量由初始值决定
	array3 := [...]int{10, 20, 30, -10, 50}
	fmt.Println(array3)

	//长度为5，指定初始化的索引
	array4 := [5]int{1: 10, 2: 30}
	fmt.Println(array4)

	//修改索引对应值
	array2[2] = 35

	//声明指针数组
	arrayIndex := [5]*int{0: new(int), 1: new(int)}
	*arrayIndex[0] = 10
	*arrayIndex[1] = 20

	var arrayString [5]string
	arrayString2 := [5]string{"Red", "Yello", "Green", "Blue"}
	arrayString = arrayString2
	fmt.Println(arrayString)

	//指针数组赋值给另一个数组
	var arrIndex1 [3]*string
	arrIndex2 := [3]*string{new(string), new(string), new(string)}
	*arrIndex2[0] = "Red"
	*arrIndex2[1] = "Blue"
	*arrIndex2[2] = "Yellow"
	arrIndex1 = arrIndex2
	fmt.Println(arrIndex1)
	fmt.Printf("value :%d", &arrIndex1)

	//二维数组
	//指定外层索引
	arr := [4][2]int{1: {20, 21}, 3: {10, -20}}
	//指定外层 和 内层 索引
	arr2 := [4][2]int{1: {0: 20}, 3: {1: 41}}

	//将arr中索引为1的维度复制到同类型中
	var arr3 [2]int = arr[1]
	//将外层索引为1、内层索引为0 复制
	var value int = arr[1][0]
	fmt.Println(arr, arr2, arr3, value)

	//大数据量使用指针进行函数传值
	var arrbig [1000000]int
	foo(&arrbig)
}

func foo(arrbig *[1000000]int) {

}
