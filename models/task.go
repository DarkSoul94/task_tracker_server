package models

import "time"

type Task struct {
	ID           uint64
	Name         string
	Description  string
	CreationDate time.Time
	InWorkTime   time.Duration
	Author       *User
	Developer    *User
	Customer     *User
	Status       *TaskStatus
	Category     *Category
	Project      *Project
	Priority     bool
	ExecOrder    int
}

func (t *Task) FillNewTask() {
	t.Status = &TaskStatus{}
	if t.Developer == nil {
		t.Status.Set(KeyTSNew)
	} else {
		t.Status.Set(KeyTSQuery)
	}
	t.CreationDate = time.Now().Truncate(time.Second)
	t.InWorkTime = 0
	if t.ExecOrder == 0 {
		t.ExecOrder = 1
	}
}
