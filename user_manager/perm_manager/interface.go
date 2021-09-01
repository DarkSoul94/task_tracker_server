package perm_manager

type PermManager interface {
	ParseFromActions(actions ...string) error
}
