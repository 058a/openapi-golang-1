package repository

import (
	"context"
	"database/sql"

	"openapi/internal/domain/entity"
	"openapi/internal/domain/model"
	"openapi/internal/infra/sqlboiler"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IStockItem interface {
	Save(model *model.StockItem) error
	Get(stockItemId entity.StockItemId) (*model.StockItem, error)
	Find(stockItemId entity.StockItemId) (bool, error)
}

type StockItem struct {
	IStockItem
	*sql.DB
}

func (r *StockItem) Save(model *model.StockItem) error {

	data := &sqlboiler.StockItem{
		ID:   model.GetId().ToUuid().String(),
		Name: model.GetName(),
	}

	err := data.Upsert(
		context.Background(),
		r.DB,
		true,
		[]string{"id"},
		boil.Whitelist("name", "deleted"),
		boil.Infer(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *StockItem) Get(stockItemId entity.StockItemId) (*model.StockItem, error) {

	id := stockItemId.ToUuid().String()
	data, err := sqlboiler.FindStockItem(context.Background(), r.DB, id)
	if err != nil {
		return &model.StockItem{}, err
	}

	model := model.NewStockItem(stockItemId.ToUuid(), data.Name)

	return model, nil
}

func (r *StockItem) Find(stockItemId entity.StockItemId) (bool, error) {
	id := uuid.UUID(stockItemId).String()
	found, err := sqlboiler.StockItemExists(context.Background(), r.DB, id)
	if err != nil {
		return false, err
	}

	return found, nil
}
