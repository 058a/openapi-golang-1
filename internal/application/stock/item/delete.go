package item

import (
	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

type DeleteRequestDto struct {
	Id uuid.UUID
}

type DeleteResponseDto struct {
}

func Delete(req *DeleteRequestDto, r item.IRepository) (*DeleteResponseDto, error) {

	id := item.Id(req.Id)
	aggregate, err := r.Get(id)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	aggregate.Delete()

	err = r.Save(aggregate)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	return &DeleteResponseDto{}, nil
}
