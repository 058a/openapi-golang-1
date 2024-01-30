package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

func TestStockItemId(t *testing.T) {

	id := uuid.New()
	stockItemId := item.Id(id)

	if stockItemId.ToUuid() != id {
		t.Errorf("expected %s, got %s", id, stockItemId.ToUuid())
	}

	if stockItemId.ToString() != id.String() {
		t.Errorf("expected %s, got %s", id.String(), stockItemId.ToString())
	}
}
