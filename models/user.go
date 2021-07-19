package models

type User struct {
	ID    uint64
	Name  string
	Group *Group
}
