package models

type Group struct {
	ID          uint64
	Name        string
	Permissions map[string][]string
}

const (
	//ключи для группы прав "Task"
	TasksGet_All     string = "task.get.all"
	TasksGet_ByAutor string = "task.get.by_author"
	TasksGet_ByDev   string = "task.get.by_developer"
	TasksCreate      string = "task.create"

	//ключи для группы прав "User"
	UserUpdate string = "user.update"

	//ключи для группы прав "Group"
	GroupUpdate string = "group.update"
	GroupCreate string = "group.create"
)

var AllPermissionsList []string = []string{
	TasksGet_All,
	TasksGet_ByAutor,
	TasksGet_ByDev,
	TasksCreate,
	UserUpdate,
	GroupUpdate,
	GroupCreate,
}

func (g *Group) ParsePermissionsByGroups(groups ...string) {

}
