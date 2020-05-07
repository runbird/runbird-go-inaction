package main

import "fmt"

//slice
func main() {
	//切片概念
	//抽象的数据结构，创建包括了只想底层数组的指针，切片访问的元素的个数和切片允许增长的个数

	//make创建 或者 字面量
	//如果只指定长度，那么容量和长度相等
	sliceString := make([]string, 5)
	//长度3，容量5
	sliceString2 := make([]string, 3, 5)
	//error 长度小于容量
	//	sliceError := make([]string, 5, 3)
	//	fmt.Println(sliceString, sliceString2, sliceError)
	fmt.Println(sliceString, sliceString2)

	//[] 中有数值是数组，无数值是slice
	slice := []string{"Red", "Blue", "Yellow"}
	slice2 := []string{99: ""} //初始化第100个元素
	fmt.Println(slice, slice2)

	//值为nil的切片
	var sliceNil []int

	//声明空切片
	// make创建
	sliceEmpty := make([]int, 0)
	// 字面量创建
	sliceEmpty2 := []int{}
	fmt.Println(sliceNil, sliceEmpty, sliceEmpty2)

	//使用切片创造切片，切片之所以叫切片就是对数组的切割
	sliceMake := []int{10, 20, 30, 40, 50}
	sliceMake2 := sliceMake[1:3]
	sliceMake3 := sliceMake[1:]
	// slice[i:j] 对于容量K的切片来说，长度 j-i; 容量 k-i
	fmt.Println("length of sliceMake2", len(sliceMake2))
	fmt.Println("cap of sliceMake2", cap(sliceMake2))
	fmt.Println("sliceMake", sliceMake)
	fmt.Println("sliceMake2", sliceMake2)
	fmt.Println("sliceMake3", sliceMake3)

	//切片是共享数据的，修改值
	sliceMake[2] = 40
	fmt.Println(sliceMake2)

	//切片索引越界异常，可通过appen操作增加长度，但是有可能改变原切片的值
	newSilce := append(sliceMake2, 60)
	fmt.Println(sliceMake, newSilce)

	//append改变长度和容量, 此时newslice2 有了一个全新的底层数组，数组容量也发生了改变
	newSlice2 := append(sliceMake, 70)
	fmt.Println(sliceMake, newSlice2) // [10 20 40 60 50] [10 20 40 60 50 70]

	//创建切片的第3个索引，进行安全性保护
	source := []string{"Red", "Blue", "Yellow", "Orange", "White"}
	//长度为1 容量为1
	//限制容量防止 共享底层数组，append修改原有数组的内容，append超过容量后新开一个底层数组
	source2 := source[2:3:3]
	fmt.Println(source, "\n", source2)
	source2 = append(source2, "wiki")
	fmt.Println(source, "\n", source2)

	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))

	//数字 %d  字符串 %v
	//range从头开始迭代
	for index, value := range source {
		fmt.Printf("source :%v value: %v\n", index, value)
	}

	rangeSlice := []int{10, 20, 30, 40, 50}
	for i, v := range rangeSlice {
		fmt.Printf("Value: %d Value-Addr: %X ElementAddr :%X\n",
			v, &v, &rangeSlice[i])
	}

	//多维切片 嵌入切片
	//mutSlice -> 0 , 1 , 2 皆是指针指向 -> 底层数组{10}, {20, 30}, {40, 50, 60}
	mutSlice := [][]int{{10}, {20, 30}, {40, 50, 60}}
	//为第一个切片追加20
	mutSlice[0] = append(mutSlice[0], 20)

	//分配一个100W个整型值的切片
	fooSlice := make([]int, 1e6)
	//64位机器上，一个切片需要24字节内容，指针需要8字节，长度和容量各需要8字节
	fooSlice = foo(fooSlice)

}

//函数间只需传递24字节的数据
func foo(fooSlice []int) []int {
	return fooSlice
}
