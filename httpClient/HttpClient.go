package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	data := make(url.Values)
	data["userName"] = []string{"abin"}
	data["passWord"] = []string{"myPassWord"}

	res, err := http.PostForm("http://localhost:9000/login", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	fmt.Println("send post success")
}
