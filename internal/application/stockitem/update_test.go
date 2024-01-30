package stockitem_test

import (
	"openapi/internal/application/stockitem"

	"openapi/internal/domain/stock/item"
	"openapi/internal/infra/database"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateSuccess(t *testing.T) {
	// Setup
	db, err := database.Open()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	repository := &item.Repository{DB: db}

	// Given
	beforeName := uuid.NewString()
	afterName := uuid.NewString()

	reqCreateDto := &stockitem.CreateRequestDto{
		Name: beforeName,
	}

	resCreateDto, err := stockitem.Create(reqCreateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// When
	reqUpdateDto := &stockitem.UpdateRequestDto{
		Id:   resCreateDto.Id,
		Name: afterName,
	}

	resUpdateDto, err := stockitem.Update(reqUpdateDto, repository)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if resUpdateDto == nil {
		t.Errorf("expected not empty, actual empty")
	}

	model, err := repository.Get(item.Id(resCreateDto.Id))
	if err != nil {
		t.Fatal(err)
	}

	if model.GetName() != afterName {
		t.Errorf("expected %s, got %s", afterName, model.GetName())
	}
}
