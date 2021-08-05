package models

type Group struct {
	ID          uint64
	Name        string
	Permissions map[string][]string
}

const (
	//ключи для группы прав "Task"
	PermGroupTasks         string = "task"
	PermGetAllTasks        string = "get_all_tasks"
	PermGetTasksWhereAutor string = "get_tasks_where_author"
	PermGetTasksWhereDev   string = "get_tasks_where_developer"
	PermCreateTask         string = "create_task"
	PermRemoveTask         string = "remove_task"

	//ключи для группы прав "User"
	PermGroupUser string = "user"
	PermUpdGroup  string = "upd_group"
	PermUpdUser   string = "upd_user"
)

var AllPermissionsList map[string][]string = map[string][]string{
	PermGroupTasks: {
		PermGetAllTasks,
		PermGetTasksWhereAutor,
		PermGetTasksWhereDev,
		PermCreateTask,
		PermRemoveTask,
	},
	PermGroupUser: {
		PermUpdGroup,
		PermUpdUser,
	},
}
