package models

import "time"

type TaskForList struct {
	ID           uint64
	Name         string
	Description  string
	CreationDate time.Time
	InWorkTime   time.Duration
	Status       *TaskStatus
	Priority     bool
	ExecOrder    uint64
}
