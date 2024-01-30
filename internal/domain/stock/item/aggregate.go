package item

import (
	"github.com/google/uuid"
)

type Aggregate struct {
	id      Id
	name    string
	deleted bool
}

func NewAggregate(id uuid.UUID, name string) *Aggregate {
	return &Aggregate{
		id:      Id(id),
		name:    name,
		deleted: false,
	}
}

func (a *Aggregate) GetId() Id {
	return a.id
}

func (a *Aggregate) GetName() string {
	return a.name
}

func (a *Aggregate) SetName(name string) {
	a.name = name
}

func (a *Aggregate) IsDeleted() bool {
	return a.deleted
}

func (a *Aggregate) Delete() {
	a.deleted = true
}
