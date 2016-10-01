package main

import (
	"fmt"
)

func createMap() {
	fmt.Println("createMap" + "-----------start")
	var smap = make(map[string]string)
	smap["a"] = "aa"
	smap["b"] = "bb"
	fmt.Println(smap)

	var cmap = map[string]string{
		"c": "cc",
		"d": "dd",
	}
	fmt.Println(cmap)
	if v, ok := cmap["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	for k, v := range cmap {
		fmt.Println(k, v)
	}

	fmt.Println("createMap" + "-----------end")
}

func main() {
	createMap()

}
