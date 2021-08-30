package models

type TaskStatus struct {
	ID   uint64
	Name string
}

const (
	KeyTSNew       string = "новая"
	KeyTSQuery     string = "очередь"
	KeyTSInWork    string = "работа"
	KeyTSPaused    string = "приостановлена"
	KeyTSWaiting   string = "ожидание"
	KeyTSRejected  string = "отклонена"
	KeyTSTesting   string = "тестирование"
	KeyTSCompleted string = "выполнена"
)

var TaskStatusMap map[string]TaskStatus = map[string]TaskStatus{
	KeyTSNew: {
		ID:   1,
		Name: "Новая",
	},
	KeyTSQuery: {
		ID:   2,
		Name: "Очередь к реализации",
	},
	KeyTSInWork: {
		ID:   3,
		Name: "В работе",
	},
	KeyTSPaused: {
		ID:   4,
		Name: "Приостановлена",
	},
	KeyTSWaiting: {
		ID:   5,
		Name: "Ожидание",
	},
	KeyTSRejected: {
		ID:   6,
		Name: "Отклонена",
	},
	KeyTSTesting: {
		ID:   7,
		Name: "Готово к тестированию",
	},
	KeyTSCompleted: {
		ID:   8,
		Name: "Выполнено",
	},
}

func (ts *TaskStatus) Set(key string) {
	var temp TaskStatus = TaskStatus{}
	temp = TaskStatusMap[key]
	*ts = temp
}
