package model

import "github.com/google/uuid"

type StockLocationId uuid.UUID

type StockLocation struct {
	Id      StockLocationId
	Name    string
	Deleted bool
}

func NewStockLocation(id StockLocationId, name string) *StockLocation {
	return &StockLocation{
		Id:      id,
		Name:    name,
		Deleted: false,
	}
}

func (s *StockLocation) Delete() {
	s.Deleted = true
}
