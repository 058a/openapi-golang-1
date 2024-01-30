package entity

import "github.com/google/uuid"

type StockItemId uuid.UUID

func (v StockItemId) ToUuid() uuid.UUID {
	return uuid.UUID(v)
}

func (v StockItemId) ToString() string {
	return uuid.UUID(v).String()
}
