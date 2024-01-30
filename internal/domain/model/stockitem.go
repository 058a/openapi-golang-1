package model

import (
	"openapi/internal/domain/entity"

	"github.com/google/uuid"
)

type StockItem struct {
	id      entity.StockItemId
	name    string
	deleted bool
}

func NewStockItem(id uuid.UUID, name string) *StockItem {
	return &StockItem{
		id:      entity.StockItemId(id),
		name:    name,
		deleted: false,
	}
}

func (s *StockItem) GetId() entity.StockItemId {
	return s.id
}

func (s *StockItem) GetName() string {
	return s.name
}

func (s *StockItem) IsDeleted() bool {
	return s.deleted
}
func (s *StockItem) SetName(name string) {
	s.name = name
}

func (s *StockItem) Delete() {
	s.deleted = true
}
