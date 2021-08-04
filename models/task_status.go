package models

type TaskStatus struct {
	ID   uint64
	Name string
}

const (
	KeyTSNew       string = "новая"
	KeyTSRevision  string = "доработка"
	KeyTSInWork    string = "работа"
	KeyTSPostponed string = "отложена"
	KeyTSCompleted string = "выполнена"
	KeyTSRejected  string = "отклонена"
)

var TaskStatusMap map[string]TaskStatus = map[string]TaskStatus{
	KeyTSNew: {
		ID:   1,
		Name: "Новая",
	},
	KeyTSRevision: {
		ID:   2,
		Name: "В доработке",
	},
	KeyTSInWork: {
		ID:   3,
		Name: "В работе",
	},
	KeyTSPostponed: {
		ID:   4,
		Name: "Отложена",
	},
	KeyTSCompleted: {
		ID:   5,
		Name: "Выполнена",
	},
	KeyTSRejected: {
		ID:   6,
		Name: "Отклонена",
	},
}

func (ts *TaskStatus) Set(key string) {
	var temp TaskStatus = TaskStatus{}
	temp = TaskStatusMap[key]
	*ts = temp
}
