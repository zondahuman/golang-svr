package model

/**
 *
 * 聊天请求实体类
 *
**/

type ChatMessageResponse struct {
	base     BaseModel
	Id       int
	UserName string
	Message  string
}
