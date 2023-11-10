package serialize

import (
	"Memorandum/model"
)

//type CreateTask struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//	Status  int    `json:"status"`
//}
//
//type UpdateTask struct {
//	Title string `json:"title"`
//}
//
//func createTask(task CreateTask)

func createTask(item model.Task) model.Task {
	return model.Task{
		Uid:       item.Uid,
		Id:        item.Id,
		Title:     item.Title,
		Content:   item.Content,
		View:      item.View,
		Status:    item.Status,
		CreateAt:  item.CreateAt,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

func CreateTasks(items []model.Task) (tasks []model.Task) {
	for _, v := range items {
		task := createTask(v)
		tasks = append(tasks, task)
	}
	return tasks
}
