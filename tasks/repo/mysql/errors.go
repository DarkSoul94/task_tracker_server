package mysql

import "errors"

var (
	ErrTaskNotExist     = errors.New("Task not exist")
	ErrFailedWriteTrack = errors.New("Failed write task track to db")
	ErrFailedReadTrack  = errors.New("Failed read task track from db")
)
