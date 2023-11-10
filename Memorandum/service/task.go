package service

import (
	"Memorandum/config"
	"Memorandum/model"
	"Memorandum/serialize"
	"time"
)

//func createTask(task serialize.CreateTask, id uint) model.Task {
//	var retTask model.Task
//	retTask = model.Task{
//		Id:       id,
//		Title:    task.Title,
//		Content:  task.Content,
//		Status:   0, //0未完成，1已完成
//		CreateAt: time.Now().Unix(),
//	}
//	return retTask
//}
//
//func UpdateTask(task serialize.UpdateTask, id uint) {
//
//}

type CreateTask struct {
	Id      uint   `json:"id"`
	Status  uint   `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateTask struct {
	Status  int    `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Id      uint   `json:"id"`
}

type DeleteTask struct {
	Id int `json:"id"`
}

type DeleteTasksInStatus struct {
	Status int `json:"status"`
}

type DeleteAllTasks struct {
}

type ShowTasksInStatus struct {
	Status int `json:"status"`
}

type ShowAllTasks struct {
}

type ShowTasksByKey struct {
	Key string `json:"key"`
}

func (ct *CreateTask) CreateTask(uid uint) serialize.Response {
	code := 200
	var count int64
	config.DB.First("id=?", ct.Id).Count(&count)
	if count != 0 {
		return serialize.Response{
			Msg:    "编号已存在！",
			Status: 500,
		}
	}
	task := model.Task{
		Uid:       uid,   //用户ID
		Id:        ct.Id, //备忘录编号
		Title:     ct.Title,
		Content:   ct.Content,
		View:      0,
		Status:    int(ct.Status),
		CreateAt:  time.Now().Unix(),
		StartTime: 0,
		EndTime:   0,
	}

	if err := config.DB.Create(&task).Error; err != nil {
		return serialize.Response{
			Status: 500,
			Msg:    "创建任务失败!",
			Error:  err,
		}
	}
	return serialize.Response{
		Status: uint(code),
		Msg:    "创建成功!",
		Data:   task,
	}
}

func (ut *UpdateTask) UpdateTask(uid uint) serialize.Response {
	var count int64
	task := model.Task{
		Uid:       uid,
		Id:        ut.Id,
		Title:     ut.Title,
		Content:   ut.Content,
		View:      0,
		Status:    ut.Status,
		CreateAt:  time.Now().Unix(),
		StartTime: 0,
		EndTime:   0,
	}
	config.DB.Model(task).Where("uid=?", uid).Where("id=?", ut.Id).Count(&count)
	if count == 0 {
		return serialize.Response{
			Msg:    "查无此事项!",
			Status: 500,
		}
	}
	config.DB.Save(task)
	return serialize.Response{
		Msg:    "更新成功!",
		Status: 200,
	}
}

func (dt *DeleteTask) DeleteTask(uid uint) serialize.Response {
	var count int64
	task := model.Task{}
	config.DB.Model(task).Where("uid=?", uid).Where("id=?", dt.Id).Count(&count)
	if count == 0 {
		return serialize.Response{
			Msg:    "所删除事项不存在！",
			Status: 500,
		}
	}
	config.DB.Where("uid=?", uid).Where("id=?", dt.Id).Delete(&task)
	return serialize.Response{
		Msg:    "事项删除成功!",
		Status: 200,
	}
}

func (dt *DeleteTasksInStatus) DeleteTasksInStatus(uid uint) serialize.Response {
	task := model.Task{}
	var count int64
	if dt.Status != 1 && dt.Status != 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "状态输入错误",
		}
	}

	config.DB.Model(task).Where("uid=?", uid).Where("status=?", dt.Status).Delete(&task).Count(&count)
	if count == 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "无可删除事项",
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "删除成功！",
	}

}

func (dt *DeleteAllTasks) DeleteAllTasks(uid uint) serialize.Response {
	task := model.Task{}
	var count int64
	config.DB.Model(task).Where("uid=?", uid).Delete(&task).Count(&count)
	if count == 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "无可删除事项",
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "删除成功！",
	}
}

func (st *ShowTasksInStatus) ShowTasksInStatus(uid uint) serialize.Response {
	var tasks []model.Task
	config.DB.Where("uid=?", uid).Where("status=?", st.Status).
		Find(&tasks)
	if len(tasks) == 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "无",
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "成功展示",
		Data:   serialize.CreateTasks(tasks),
	}
}

func (st *ShowAllTasks) ShowAllTasks(uid uint) serialize.Response {
	var task []model.Task
	var tasks []model.Task
	config.DB.Model(&task).Where("uid=?", uid).
		Find(&tasks)
	if len(tasks) == 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "无",
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "成功展示",
		Data:   serialize.CreateTasks(tasks),
	}
}

func (st *ShowTasksByKey) ShowTasksByKey(uid int) serialize.Response {
	var task []model.Task
	var tasks []model.Task
	config.DB.Model(&task).Where("uid=?", uid).Where("content LIKE ?", "%"+st.Key+"%").
		Find(&tasks)
	if len(tasks) == 0 {
		return serialize.Response{
			Status: 500,
			Msg:    "无",
		}
	}
	return serialize.Response{
		Status: 200,
		Msg:    "成功展示",
		Data:   serialize.CreateTasks(tasks),
	}
}
