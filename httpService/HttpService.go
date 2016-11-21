package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(response http.ResponseWriter, request *http.Request) {
		if request.Method == "POST" {
			var (
				userName string = request.PostFormValue("userName")
				passWord string = request.PostFormValue("passWord")
			)
			fmt.Println("userName is : %s ", userName)
			fmt.Println("passWord is : %s ", passWord)
		}
		response.Write([]byte("login success"))
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
