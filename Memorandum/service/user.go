package service

import (
	"Memorandum/config"
	"Memorandum/model"
	"Memorandum/serialize"
	"github.com/cloudwego/hertz/pkg/app"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *app.RequestContext) serialize.Response {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return serialize.Response{
			Error:  err,
			Msg:    "输入不合规定",
			Status: 400,
		}
	}
	var count int64
	config.DB.Where("id=?", user.Id).First(user).Count(&count)
	if count != 0 {

		return serialize.Response{
			Msg:    "账号已存在!请重新输入",
			Status: 400,
		}
	} else {
		//应该先对密码进行加密处理，然后再进行存储
		passWord, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.CreateUser(passWord) //进行加密处理
		//myToken, _ := generateToken(user.Id, user.UserName, user.Password)
		config.DB.Create(&user)
		return serialize.Response{
			Msg:    "注册成功!",
			Status: 200,
		}
	}
}

func Login(ctx *app.RequestContext) serialize.Response {
	var user model.User
	var userTemp model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		return serialize.Response{
			Error:  err,
			Msg:    "输入不合规定",
			Status: 400,
		}
	}
	var count int64
	config.DB.Where("id=?", user.Id).First(&userTemp).Count(&count)
	if count == 0 {
		return serialize.Response{
			Msg:    "查无此号",
			Status: 400,
		}
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(userTemp.Password), []byte(user.Password))
		if err == nil {
			//验证成功了

			myToken, _ := GenerateToken(userTemp.Id, userTemp.UserName, userTemp.Password)
			//dataToken
			return serialize.Response{
				Msg:    "登录成功",
				Status: 200,
				Data:   serialize.TokenData{Id: user.Id, Data: myToken},
			}
		} else {
			return serialize.Response{
				Msg:    "验证失败，发生错误",
				Status: 400,
			}
		}
	}
}
