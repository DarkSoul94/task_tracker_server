package usecase

import (
	"github.com/DarkSoul94/task_tracker_server/models"
	"github.com/DarkSoul94/task_tracker_server/tasks"
	"github.com/DarkSoul94/task_tracker_server/user_manager"
)

var (
	checksList map[string]func(map[string][]string) string = map[string]func(map[string][]string) string{
		tasks.KeyGet:           getTasks,
		user_manager.KeyGet:    getUsersList,
		user_manager.KeyUpdate: updateUser,
	}
)

func getTasks(permissions map[string][]string) string {
	if contains(permissions[tasks.KeyGet], models.TasksGet_All) {
		return tasks.KeyGet_All
	}
	if contains(permissions[tasks.KeyGet], models.TasksGet_ByAutor, models.TasksGet_ByDev) {
		return tasks.KeyGet_AuthorDev
	}
	if contains(permissions[tasks.KeyGet], models.TasksGet_ByAutor) {
		return tasks.KeyGet_Author
	}
	if contains(permissions[tasks.KeyGet], models.TasksGet_ByDev) {
		return tasks.KeyGet_Dev
	}
	return ""
}

func getUsersList(permissions map[string][]string) string {
	if !contains(permissions[user_manager.KeyGet], models.UserGet) {
		return ""
	}
	return user_manager.KeyGet
}

func updateUser(permissions map[string][]string) string {
	if !contains(permissions[user_manager.KeyUpdate], models.UserUpdate) {
		return ""
	}
	return user_manager.KeyUpdate
}

func contains(slice []string, targets ...string) (result bool) {
	temp := make([]string, 0)
	for _, val := range slice {
		for _, target := range targets {
			if target == val {
				temp = append(temp, target)
			}
		}
	}
	if len(temp) == len(targets) {
		result = true
	}
	return result
}
