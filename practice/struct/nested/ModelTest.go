package main

import (
	"fmt"
	"golang-svr/practice/struct/model"
)

/**
 *
 * 基本状态基类
 *
**/

func main() {
	chatMessageResponse := model.ChatMessageResponse{base: model.BaseModel{Status: "SUCCESS", Message: "no message", Data: nil}, Id: 5, UserName: "abin", Message: "hi"}
	fmt.Println("chatMessageResponse", chatMessageResponse.Id)
}
