package model

type Task struct {
	Uid       uint   `json:"uid"`
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      uint   `json:"view"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}
