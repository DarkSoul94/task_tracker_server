package permissions

type PermManager interface {
	ExportToTreeView() []byte
}
