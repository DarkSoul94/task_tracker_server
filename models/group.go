package models

type Group struct {
	ID          uint64
	Name        string
	Permissions map[string][]string
}
