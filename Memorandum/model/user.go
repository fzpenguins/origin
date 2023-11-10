package model

type User struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (user *User) CreateUser(passWord []byte) {
	user.Password = string(passWord)
}
