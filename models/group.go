package models

import (
	"fmt"
	"regexp"
	"strings"
)

type Group struct {
	ID          uint64
	Name        string
	Permissions map[string][]string
}

const (
	//ключи для группы прав "Task"
	TasksGet_All     string = "task.get.all"
	TasksGet_ByAutor string = "task.get.author"
	TasksGet_ByDev   string = "task.get.developer"
	TasksCreate      string = "task.create"

	//ключи для группы прав "User"
	UserUpdate string = "user.update"
	UserGet    string = "user.get"

	//ключи для группы прав "Group"
	GroupUpdate string = "group.update"
	GroupCreate string = "group.create"
)

var AllPermissionsList map[string][]string = map[string][]string{
	"task": {
		TasksGet_All,
		TasksGet_ByAutor,
		TasksGet_ByDev,
		TasksCreate,
	},
	"user": {
		UserUpdate,
		UserGet,
	},
	"group": {
		GroupUpdate,
		GroupCreate,
	},
}

func (g *Group) ParsePermissionsByAction(actions ...string) map[string][]string {
	permissions := make(map[string][]string)
	for _, key := range actions {
		subKeys := strings.Split(key, ".")

		format := g.prepActionKeyToRegExp(subKeys)
		fmt.Println(format)
		re := regexp.MustCompile(format)
		for _, val := range g.Permissions[subKeys[0]] {
			if re.MatchString(val) {
				permissions[key] = append(permissions[key], val)
			}
		}
	}
	return permissions
}

func (g *Group) prepActionKeyToRegExp(subKeys []string) string {
	var result string
	if len(subKeys) > 1 {
		for i, val := range subKeys {
			if i == 0 {
				result = fmt.Sprintf("^%s", val)
			} else {
				result = fmt.Sprintf("%s[.]%s", result, val)
			}
		}
	} else {
		return subKeys[0]
	}

	return result
}
