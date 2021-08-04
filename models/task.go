package models

import "time"

type Task struct {
	ID          uint64
	Name        string
	Description string
	CreateDate  time.Time
	InWorkTime  time.Duration
	Author      *User
	Developer   *User
	Status      *TaskStatus
}

func (t *Task) FillNewTask() {
	temp := TaskStatus{}
	temp.Set(KeyTSNew)
	t.Status = &temp
	t.CreateDate = time.Now().Truncate(time.Second)
	t.InWorkTime = 0
}
