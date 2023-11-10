package serialize

import "time"

type userInfo struct {
	Id                uint   `json:"id"`
	UserName          string `json:"user_name"`
	EncryptedPassword string `json:"Encrypted_password"`
	CreateAt          int64  `json:"create_at"`
}

func (UserInfo *userInfo) createUserInfo(password []byte) {
	UserInfo.EncryptedPassword = string(password)
	UserInfo.CreateAt = time.Now().Unix()
}
