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
	t.Status = &TaskStatus{}
	t.Status.Set(KeyTSNew)
	t.CreateDate = time.Now().Truncate(time.Second)
	t.InWorkTime = 0
}
