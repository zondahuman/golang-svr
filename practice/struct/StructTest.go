package main

import (
	"fmt"
)

func createStruct() {
	type Person struct {
		name string
		age  int
	}

	var person1 Person
	person1.name = "abin"
	person1.age = 22
	fmt.Println(person1)

	person2 := Person{"steven", 21}
	fmt.Println(person2)

	person3 := Person{age: 24, name: "john"}
	fmt.Println(person3)
}

func createManyStruct() {
	type Human struct {
		name string
		age  int
	}
	type Student struct {
		Human
		speciality string
	}

	mark := Student{Human{"bill", 31}, "Computer Science"}
	fmt.Println(mark)
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His speciality is ", mark.speciality)
	fmt.Println("His mark.Human.age is ", mark.Human.age)
	fmt.Println("His mark.Human.name is ", mark.Human.name)
}

func main() {
	createStruct()
	createManyStruct()
}
