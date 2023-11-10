package api

import (
	"Memorandum/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Register(c context.Context, ctx *app.RequestContext) {
	res := service.Register(ctx)
	if res.Status != 200 {
		ctx.JSON(int(res.Status), res)
	} else {
		ctx.JSON(int(res.Status), res)
	}

}

func Login(c context.Context, ctx *app.RequestContext) {
	res := service.Login(ctx)
	if res.Status != 200 {
		ctx.JSON(int(res.Status), res)
	} else {
		ctx.JSON(int(res.Status), res)
	}
}
