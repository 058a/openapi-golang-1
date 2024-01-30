package item_test

import (
	"context"
	"openapi/internal/domain/stock/item"
	"openapi/internal/infra/database"
	"openapi/internal/infra/sqlboiler"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestStockItemSave(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	r := item.Repository{Db: db}

	// Given
	id := uuid.New()
	name := "test"
	model := item.NewModel(id, name)

	// When
	timeFormat := "2006/01/02 15:04:05.000000"
	currentDateTime := time.Now().UTC().Format(timeFormat)
	err = r.Save(model)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	data, err := sqlboiler.FindStockItem(context.Background(), db, id.String())
	if err != nil {
		t.Fatal(err)
	}

	if data.ID != id.String() {
		t.Errorf("expected %s, got %s", id, data.ID)
	}
	if data.Name != name {
		t.Errorf("expected %s, got %s", name, data.Name)
	}
	if data.Deleted != false {
		t.Errorf("expected %t, got %t", false, data.Deleted)
	}
	if data.CreatedAt.Format(timeFormat) < currentDateTime {
		t.Errorf("expected %s, got %s", currentDateTime, data.CreatedAt)
	}
	if data.UpdatedAt.Format(timeFormat) < currentDateTime {
		t.Errorf("expected %s, got %s", currentDateTime, data.UpdatedAt)
	}
	if data.UpdatedAt.Format(timeFormat) != data.CreatedAt.Format(timeFormat) {
		t.Errorf("expected %s, got %s", data.CreatedAt, data.UpdatedAt)
	}

}
