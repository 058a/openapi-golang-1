package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type UpdateRequestDto struct {
	Id   uuid.UUID
	Name string
}

type UpdateResponseDto struct{}

func Update(req *UpdateRequestDto, r item.IRepository) (*UpdateResponseDto, error) {

	id := item.Id(req.Id)
	aggregate, err := r.Get(id)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	aggregate.SetName(req.Name)

	err = r.Save(aggregate)
	if err != nil {
		return &UpdateResponseDto{}, err
	}

	return &UpdateResponseDto{}, nil
}
