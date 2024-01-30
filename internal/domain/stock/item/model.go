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

func (s *Model) GetId() Id {
	return s.id
}

func (s *Model) GetName() string {
	return s.name
}

func (s *Model) IsDeleted() bool {
	return s.deleted
}
func (s *Model) SetName(name string) {
	s.name = name
}

func (s *Model) Delete() {
	s.deleted = true
}
