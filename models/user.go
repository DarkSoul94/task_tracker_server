package models

type User struct {
	ID       uint64
	Name     string
	PassHash string
	Group    *Group
}
