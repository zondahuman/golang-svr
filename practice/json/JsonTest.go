package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name    string
	Age     int
	Guake   bool
	Classes []string
	Price   float32
}

func createJson() string {
	fmt.Println("createJson", "---------start")
	stud := &Student{
		"XiaoMing",
		16,
		true,
		[]string{"Math", "English", "Chinese"},
		9.99,
	}

	studJson, err := json.Marshal(stud)
	if err != nil {
		fmt.Println(err.Error())
	}
	var result = string(studJson)
	fmt.Println(result)
	fmt.Println("createJson", "---------end")
	return result
}

func parseJson() {
	fmt.Println("parseJson", "---------init")
	var studJson = createJson()
	fmt.Println("parseJson", "---------start")
	stud := &Student{}

	json.Unmarshal([]byte(studJson), &stud)
	fmt.Println(stud)
	fmt.Println("parseJson", "---------end")
}

func main() {
	createJson()
	parseJson()
}
