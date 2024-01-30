package model_test

import (
	"openapi/internal/domain/model"
	"testing"

	"github.com/google/uuid"
)

func TestStockLocationNew(t *testing.T) {
	// When
	generatedUuid := uuid.New()
	id := model.StockLocationId(generatedUuid)
	name := "test"
	stockLocation := model.NewStockLocation(id, name)

	// Then
	if stockLocation.Id != id {
		t.Errorf("expected %s, got %s", id, stockLocation.Id)
	}
	castedId := uuid.UUID(stockLocation.Id)
	if castedId != generatedUuid {
		t.Errorf("expected %s, got %s", generatedUuid, castedId)
	}
	if stockLocation.Name != name {
		t.Errorf("expected %s, got %s", name, stockLocation.Name)
	}
}

func TestStockLocationDelete(t *testing.T) {
	// Given
	generatedUuid := uuid.New()
	id := model.StockLocationId(generatedUuid)
	name := "test"
	stockLocation := model.NewStockLocation(id, name)

	// When
	stockLocation.Delete()

	// Then
	if stockLocation.Deleted != true {
		t.Errorf("expected %t, got %t", true, stockLocation.Deleted)
	}
}
