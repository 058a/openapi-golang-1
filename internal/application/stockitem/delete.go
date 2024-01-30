package stockitem

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
	model, err := r.Get(id)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	model.Delete()

	err = r.Save(model)
	if err != nil {
		return &DeleteResponseDto{}, err
	}

	return &DeleteResponseDto{}, nil
}
