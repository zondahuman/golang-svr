package main

import (
	"fmt"
)

func stringArrayTest() {
	var sarray = []string{"abin", "lee", "john"}
	fmt.Println(sarray)
	sarray = append(sarray, "how")
	fmt.Println(sarray)
}

func intArrayTest() {
	var iarray = []int{1, 2, 3}
	fmt.Println(iarray)
	iarray = append(iarray, 5)
	fmt.Println(iarray)
	var marray = make([]int, 5, 10)
	fmt.Println(marray)
	marray[0] = 1
	fmt.Println(marray)
}

func loopIntArray() {
	fmt.Println("loopIntArray" + "---------start")
	var iarray = []int{1, 2, 3}
	fmt.Println(iarray)
	//传统方式：
	for i := 0; i < len(iarray); i++ {
		fmt.Println("common-iarray[", i, "]=", iarray[i])
	}
	//使用range表达式
	//range表达式有两个返回值，第一个是索引，第二个是元素的值：
	for i, v := range iarray {
		fmt.Println("range---iarray[", i, "]=", v)
	}

	fmt.Println("loopIntArray", "---------end")
}

func main() {
	stringArrayTest()
	intArrayTest()
	loopIntArray()
}
