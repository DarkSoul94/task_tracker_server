package http

import "github.com/DarkSoul94/task_tracker_server/auth"

// Handler ...
type Handler struct {
	ucAuth auth.AuthUC
}

type Response struct {
	Status string      `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data"`
}

type inpGroup struct {
	ID          uint64              `json:"group_id"`
	Name        string              `json:"group_name"`
	Permissions map[string][]string `json:"permissions"`
}

type loginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//ValidateData проверяет авторизационные данные пользователя на соответствие установленным правилам.
func (l *loginUser) ValidateData() error {
	if !l.validateUserName() {
		return ErrUserNameRequirements
	}
	if !l.validatePassword() {
		return ErrPassRequirements
	}
	return nil
}

func (l *loginUser) validateUserName() bool {
	for _, val := range userNameReq {
		if !val(&l.UserName) {
			return false
		}
	}
	return true
}

func (l *loginUser) validatePassword() bool {
	for _, val := range passwordReq {
		if !val(&l.Password) {
			return false
		}
	}
	return true
}

type inpUser struct {
	ID    uint64    `json:"user_id,omitempty"`
	Name  string    `json:"user_name,omitempty"`
	Token string    `json:"token"`
	Group *inpGroup `json:"group,omitempty"`
}
