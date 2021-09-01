package global_const

const (
	ActionTaskGet          = "task.get" //получение списка задач
	ActionTaskGet_author   = "task.get.author"
	ActionTaskGet_dev      = "task.get.dev"
	ActionTaskGet_customer = "task.get.customer"

	ActionTaskCreate = "task.create" //создание новой задачи

	ActionTaskUpdate = "task.update" //обновление задачи

	ActionUserUpdate = "user.update" //обновление пользователя

	ActionGroupUpdate = "group.update" //обновление группы пользователей
	ActionGroupCreate = "group.create" //создание группы пользователей
)

var ActionsForPerm []string = []string{
	ActionTaskGet_author,
	ActionTaskGet_dev,
	ActionTaskGet_customer,
	ActionTaskCreate,
	ActionTaskUpdate,
	ActionUserUpdate,
	ActionGroupUpdate,
	ActionGroupCreate,
}
