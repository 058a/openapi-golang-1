package stockitem_test

import (
	"openapi/internal/application/stockitem"
	"openapi/internal/domain/stock/item"
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

func TestCreateSuccess(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &item.Repository{DB: db}

	// Given
	reqDto := &stockitem.CreateRequestDto{
		Name: uuid.NewString(),
	}

	// When
	resDto, err := stockitem.Create(reqDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resDto.Id == uuid.Nil {
		t.Errorf("expected %s, got %s", uuid.Nil, resDto.Id)
	}

	model, err := repository.Get(item.Id(resDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if model.GetName() != reqDto.Name {
		t.Errorf("expected %s, got %s", reqDto.Name, model.GetName())
	}
}
