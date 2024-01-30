package item

import (
	"context"
	"database/sql"

	"openapi/internal/infra/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(model *Model) error
	Get(id Id) (*Model, error)
	Find(id Id) (bool, error)
}

type Repository struct {
	IRepository
	*sql.DB
}

func (r *Repository) Save(model *Model) error {

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

func (r *Repository) Get(id Id) (*Model, error) {
	data, err := sqlboiler.FindStockItem(context.Background(), r.DB, id.ToString())
	if err != nil {
		return &Model{}, err
	}

	model := NewModel(id.ToUuid(), data.Name)

	return model, nil
}

func (r *Repository) Find(id Id) (bool, error) {
	found, err := sqlboiler.StockItemExists(context.Background(), r.DB, id.ToString())
	if err != nil {
		return false, err
	}

	return found, nil
}
