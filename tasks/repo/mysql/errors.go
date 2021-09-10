package mysql

import "errors"

var (
	ErrTaskNotExist    = errors.New("Task not exist")
	ErrAddTrack        = errors.New("Failed write task track to db")
	ErrGetTrack        = errors.New("Failed read task track from db")
	ErrGetCategoryList = errors.New("Failed read categories list from db")
	ErrUpdateTask      = errors.New("Failed update task in db")
)
