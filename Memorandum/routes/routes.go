package routes

import (
	"Memorandum/api"
	"Memorandum/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

//不能用gin

func NewRouter() *server.Hertz {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	store := cookie.NewStore([]byte("something-very-secret"))

	h.Use(sessions.New("mysession", store))
	userGroup := h.Group("/api/v1")
	{
		userGroup.POST("user/register", api.Register)
		userGroup.POST("user/login", api.Login)
		task := userGroup.Group("/")
		task.Use(middleware.JWT)
		{
			task.POST("task", api.CreateTask)
			task.PUT("task/:id", api.UpdateTask)
			task.DELETE("task/:id", api.DeleteTask)
			task.DELETE("tasks/:status", api.DeleteTasksInStatus)
			task.DELETE("tasks", api.DeleteAllTasks)
			task.GET("tasks?status=:status", api.ShowTasksInStatus)
			task.GET("tasks?key=:key", api.ShowTasksByKey)
		}
	}
	return h
}
