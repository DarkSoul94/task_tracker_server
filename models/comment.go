package models

import "time"

type Comment struct {
	ID         uint64
	Text       string
	CreateTime time.Time
	Author     *User
	TaskID     uint64
}
