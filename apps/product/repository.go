package product

import (
	"context"
	"database/sql"
	"onlineShop/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}


func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreateProduct(ctx context.Context, model Product) (err error) {
	query := `
		INSERT INTO products(
			sku, name , stock, price, created_at, updated_at
		) VALUES(
		 	:sku, :name, :stock, :price, :created_at, :updated_at
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) {
	query := `
		SELECT
			id, sku, name, 
			stock, price, 
			created_at, 
			updated_at
		FROM products
		WHERE id>$1
		ORDER BY id ASC 
		LIMIT $2
	`
	err = r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return
	}
	return
}

// GetProudctBySKU implements Repository.
func (r repository) GetProudctBySKU(ctx context.Context, sku string) (product Product, err error) {
	query := `
		SELECT
			id, sku, name, 
			stock, price, 
			created_at, 
			updated_at
		FROM products
		WHERE sku=$1
	`
	r.db.GetContext(ctx, &product, query, sku)
	if err != nil {
		if err == sql.ErrNoRows{
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}