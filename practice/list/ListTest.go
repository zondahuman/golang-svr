package main

import (
	"container/list"
	"fmt"
)

func createList() {
	fmt.Println("createList" + "-----------start")
	list1 := list.New()
	for i := 0; i < 5; i++ {
		list1.PushBack(i)
	}
	fmt.Println(list1.Init())
	fmt.Println("createList" + "-----------end")
}

func loopList() {
	fmt.Println("loopList" + "-----------start")
	list1 := list.New()
	for i := 0; i < 5; i++ {
		list1.PushBack(i)
	}
	fmt.Println(list1.Back())
	fmt.Println(list1.Front())
	fmt.Println(list1.Len())

	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("loopList" + "-----------end")
}

func main() {
	createList()
	loopList()

}
