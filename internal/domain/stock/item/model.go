package item

import (
	"github.com/google/uuid"
)

type Model struct {
	id      Id
	name    string
	deleted bool
}

func NewModel(id uuid.UUID, name string) *Model {
	return &Model{
		id:      Id(id),
		name:    name,
		deleted: false,
	}
}

func (m *Model) GetId() Id {
	return m.id
}

func (m *Model) GetName() string {
	return m.name
}

func (m *Model) IsDeleted() bool {
	return m.deleted
}
func (m *Model) SetName(name string) {
	m.name = name
}

func (m *Model) Delete() {
	m.deleted = true
}
