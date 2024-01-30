package stockitem

import (
	"github.com/google/uuid"

	"openapi/internal/domain/model"
	"openapi/internal/domain/repository"
)

type CreateRequestDto struct {
	Name string
}

type CreateResponseDto struct {
	Id uuid.UUID
}

func Create(req *CreateRequestDto, r repository.IStockItem) (*CreateResponseDto, error) {

	model := model.NewStockItem(uuid.New(), req.Name)

	err := r.Save(model)
	if err != nil {
		return nil, err
	}

	return &CreateResponseDto{
		Id: model.GetId().ToUuid(),
	}, nil
}
