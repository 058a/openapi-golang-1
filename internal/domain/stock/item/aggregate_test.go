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
	aggregate := item.NewAggregate(id, name)

	// Then
	if aggregate.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, aggregate.GetId().ToUuid())
	}
	if aggregate.GetId().ToUuid() != id {
		t.Errorf("expected %s, got %s", id, aggregate.GetId().ToUuid())
	}
	if aggregate.GetName() != name {
		t.Errorf("expected %s, got %s", name, aggregate.GetName())
	}
	if aggregate.IsDeleted() != false {
		t.Errorf("expected %t, got %t", false, aggregate.IsDeleted())
	}
}

func TestMethodForDelete(t *testing.T) {
	// Given
	id := uuid.New()
	name := "test"
	stockItem := item.NewAggregate(id, name)

	// When
	stockItem.Delete()

	// Then
	if stockItem.IsDeleted() != true {
		t.Errorf("expected %t, got %t", true, stockItem.IsDeleted())
	}
}
