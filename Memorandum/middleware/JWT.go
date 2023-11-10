package middleware

import (
	"Memorandum/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"time"
)

func JWT(c context.Context, ctx *app.RequestContext) {
	var code int
	myToken := ctx.GetHeader("Authorization")
	if string(myToken) == "" {
		code = 400
		ctx.JSON(code, utils.H{
			"Msg":  "身份验证无效",
			"code": code,
		})
		ctx.Abort()
		return
	}
	claim, err := service.ParseToken(string(myToken))
	if err != nil {
		code = 400
		ctx.JSON(code, utils.H{
			"Msg":    "解析失败",
			"Status": code,
		})
		ctx.Abort()
		return
	} else {
		if claim.ExpiresAt < time.Now().Unix() {
			code = 400
			ctx.JSON(code, utils.H{
				"Msg":    "身份过期",
				"Status": code,
			})
			ctx.Abort()
			return
		}

	}
	ctx.Next(c)
}
