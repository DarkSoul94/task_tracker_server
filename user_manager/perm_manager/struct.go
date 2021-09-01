package perm_manager

type PermLayer struct {
	FinalPerm     []string
	SubPermGroups map[string]PermLayer
}
