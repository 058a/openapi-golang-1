package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestStockItemNew(t *testing.T) {
	// When
	id := uuid.New()
	name := "test"
	stockItem := item.NewModel(id, name)

	// Then
	if stockItem.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, stockItem.GetId().ToUuid())
	}
	castedId := stockItem.GetId().ToUuid()
	if castedId != id {
		t.Errorf("expected %s, got %s", id, castedId)
	}
	if stockItem.GetName() != name {
		t.Errorf("expected %s, got %s", name, stockItem.GetName())
	}
}

func TestStockItemDelete(t *testing.T) {
	// Given
	id := uuid.New()
	name := "test"
	stockItem := item.NewModel(id, name)

	// When
	stockItem.Delete()

	// Then
	if stockItem.IsDeleted() != true {
		t.Errorf("expected %t, got %t", true, stockItem.IsDeleted())
	}
}
