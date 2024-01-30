package item

import (
	"context"
	"database/sql"
	"openapi/internal/infrastructure/sqlboiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IRepository interface {
	Save(m *Model) error
	Get(id Id) (*Model, error)
	Find(id Id) (bool, error)
}

type Repository struct {
	IRepository
	Db *sql.DB
}

func (r *Repository) Save(m *Model) error {

	data := &sqlboiler.StockItem{
		ID:   m.GetId().ToUuid().String(),
		Name: m.GetName(),
	}

	err := data.Upsert(
		context.Background(),
		r.Db,
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
	data, err := sqlboiler.FindStockItem(context.Background(), r.Db, id.ToString())
	if err != nil {
		return &Model{}, err
	}

	model := &Model{
		id:      id,
		name:    data.Name,
		deleted: data.Deleted,
	}
	return model, nil
}

func (r *Repository) Find(id Id) (bool, error) {
	found, err := sqlboiler.StockItemExists(context.Background(), r.Db, id.ToString())
	if err != nil {
		return false, err
	}

	return found, nil
}
