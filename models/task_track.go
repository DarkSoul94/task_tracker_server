package models

import "time"

type TaskTrack struct {
	ID         uint64
	TaskID     uint64
	Author     *User
	StartTime  time.Time
	EndTime    time.Time
	Difference uint64
}

func (t *TaskTrack) StartTrack(taskID uint64, user *User) {
	t.ID = 0
	t.TaskID = taskID
	t.Author = user
	t.StartTime = time.Now().Truncate(time.Second)
	t.EndTime = time.Time{}
	t.Difference = 0
}

func (t *TaskTrack) EndTrack() {
	t.EndTime = time.Now().Truncate(time.Second)
	t.Difference = uint64(time.Since(t.StartTime).Seconds())
}
