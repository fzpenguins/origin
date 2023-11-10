package api

import (
	"Memorandum/serialize"
	"Memorandum/service"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateTask(c context.Context, ctx *app.RequestContext) {
	var ct service.CreateTask
	if err := ctx.BindJSON(&ct); err != nil {
		ctx.JSON(500, "参数输入错误")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))

	res := ct.CreateTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func UpdateTask(c context.Context, ctx *app.RequestContext) {
	var ut service.UpdateTask
	if err := ctx.BindJSON(&ut); err != nil {
		ctx.JSON(500, "参数输入错误")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := ut.UpdateTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func DeleteTask(c context.Context, ctx *app.RequestContext) {
	var err error
	var dt service.DeleteTask
	dt.Id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, "id参数输入有误!")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteTask(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func DeleteTasksInStatus(c context.Context, ctx *app.RequestContext) {
	var err error
	var dt service.DeleteTasksInStatus
	dt.Status, err = strconv.Atoi(ctx.Param("status"))
	if err != nil {
		ctx.JSON(500, "status参数输入有误!")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteTasksInStatus(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func DeleteAllTasks(c context.Context, ctx *app.RequestContext) {
	var dt service.DeleteAllTasks
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := dt.DeleteAllTasks(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func ShowTasksInStatus(c context.Context, ctx *app.RequestContext) {
	var st service.ShowTasksInStatus
	var sta service.ShowAllTasks
	var err error
	var res serialize.Response
	st.Status, err = strconv.Atoi(ctx.Query("status"))
	if err != nil {
		ctx.JSON(500, "status参数输入错误")
		return
	}
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	if st.Status != 0 && st.Status != 1 {
		res = sta.ShowAllTasks(claim.Id)
	} else {
		res = st.ShowTasksInStatus(claim.Id)
	}
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)

}

func ShowAllTasks(c context.Context, ctx *app.RequestContext) {
	var st service.ShowAllTasks
	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := st.ShowAllTasks(claim.Id)
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}

func ShowTasksByKey(c context.Context, ctx *app.RequestContext) {
	var st service.ShowTasksByKey
	st.Key = ctx.Query("key")

	claim, _ := service.ParseToken(string(ctx.GetHeader("Authorization")))
	res := st.ShowTasksByKey(int(claim.Id))
	if res.Status != 200 {
		ctx.JSON(500, res)
		return
	}
	ctx.JSON(200, res)
}
