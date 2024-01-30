package item

import (
	"github.com/google/uuid"

	"openapi/internal/domain/stock/item"
)

type CreateRequestDto struct {
	Name string
}

type CreateResponseDto struct {
	Id uuid.UUID
}

func Create(req *CreateRequestDto, r item.IRepository) (*CreateResponseDto, error) {

	model := item.NewModel(uuid.New(), req.Name)

	err := r.Save(model)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: model.GetId().ToUuid(),
	}, nil
}
