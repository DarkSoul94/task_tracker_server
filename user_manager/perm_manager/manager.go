package perm_manager

import (
	"fmt"
	"strings"
)

type Manager struct {
	permissions PermLayer
}

func (m *Manager) CreateManagerFromActions(actions ...string) {
	m.permissions = PermLayer{
		FinalPerm:     make([]string, 0),
		SubPermGroups: make(map[string]PermLayer),
	}
	for _, action := range actions {
		permParts := strings.Split(action, ".")
		if len(permParts) == 1 {
			m.permissions.FinalPerm = append(m.permissions.FinalPerm, permParts[0])
		} else {
			//	m.permissions.SubPermGroups[permParts[0]] = m.formPermLayer(permParts)
			m.putToLayer(&m.permissions, permParts)
		}
	}
	fmt.Println(m.permissions)
}

//TODO записывать новый слой в существующую мапу мапу
func (m *Manager) formPermLayer(actions []string) PermLayer {
	var layer PermLayer = PermLayer{
		FinalPerm:     make([]string, 0),
		SubPermGroups: make(map[string]PermLayer),
	}

	if len(actions) != 1 {
		layer.SubPermGroups[actions[0]] = m.formPermLayer(actions[1:])
	} else {
		layer.FinalPerm = append(layer.FinalPerm, actions[0])
	}

	return layer
}

func (m *Manager) putToLayer(targetLayer *PermLayer, actions []string) {
	if len(actions) != 1 {
		if _, ok := targetLayer.SubPermGroups[actions[0]]; !ok {
			targetLayer.SubPermGroups[actions[0]] = m.formPermLayer(actions[1:])
		} else {
			pointer := targetLayer.SubPermGroups[actions[0]]
			m.putToLayer(&pointer, actions[1:])
			targetLayer.SubPermGroups[actions[0]] = pointer
		}
	} else {
		targetLayer.FinalPerm = append(targetLayer.FinalPerm, actions[0])
	}
}
