package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewModel(t *testing.T) {
	// When
	id := uuid.New()
	name := "test"
	model := item.NewModel(id, name)

	// Then
	if model.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, model.GetId().ToUuid())
	}
	if model.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, model.GetId().ToUuid())
	}
	if model.GetName() != name {
		t.Errorf("expected %s, got %s", name, model.GetName())
	}
}

func TestMethodForDelete(t *testing.T) {
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
