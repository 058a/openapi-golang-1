package item_test

import (
	"openapi/internal/domain/inbound/item"
	"testing"

	"github.com/google/uuid"
)

func TestStockItemNew(t *testing.T) {
	// When
	id := uuid.New()
	name := "test"
	model := item.NewModel(id, name)

	// Then
	if model.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, model.GetId().ToUuid())
	}
	castedId := model.GetId().ToUuid()
	if castedId != id {
		t.Errorf("expected %s, got %s", id, castedId)
	}
	if model.GetName() != name {
		t.Errorf("expected %s, got %s", name, model.GetName())
	}
}

func TestStockItemDelete(t *testing.T) {
	// Given
	id := uuid.New()
	name := "test"
	model := item.NewModel(id, name)

	// When
	model.Delete()

	// Then
	if model.IsDeleted() != true {
		t.Errorf("expected %t, got %t", true, model.IsDeleted())
	}
}
