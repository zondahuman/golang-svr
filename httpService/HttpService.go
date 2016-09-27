package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			var (
				userName string = r.PostFormValue("userName")
				passWord string = r.PostFormValue("passWord")
			)
			fmt.Println("userName is : %s ", userName)
			fmt.Println("passWord is : %s ", passWord)
		}
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
