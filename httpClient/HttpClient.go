package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	data := make(url.Values)
	data["userName"] = []string{"abin"}
	data["passWord"] = []string{"myPassWord"}

	response, err := http.PostForm("http://localhost:9000/login", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	defer response.Body.Close()
	fmt.Println("send post success")
}
