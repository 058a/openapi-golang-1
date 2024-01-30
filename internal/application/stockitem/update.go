package stockitem

import (
	"openapi/internal/domain/entity"
	"openapi/internal/domain/repository"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string
}

type UpdateResponseDto struct{}

func Update(req *UpdateRequestDto, r repository.IStockItem) (*UpdateResponseDto, error) {

	id := entity.StockItemId(req.Id)
	model, err := r.Get(id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	model.SetName(req.Name)

	err = r.Save(model)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	return &UpdateResponseDto{}, nil
}
