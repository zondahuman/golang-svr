package model

/**
 *
 * 聊天请求实体类
 *
**/

type ChatMessageRequest struct {
	Id          int
	UserName    string //用户名
	Message     string //密码
	EquipmentId string //设备号
	channel     string //渠道IOS，ANDROID，PAD
	sessionId   string //
}
