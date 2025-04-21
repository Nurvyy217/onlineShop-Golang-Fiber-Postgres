package product

import (
	"context"
	"onlineShop/infra/response"
	"onlineShop/internal/log"
)

type Repository interface {
	CreateProduct(ctx context.Context, model Product) (err error)
	GetAllProductWithPaginationCursor(ctx context.Context, model ProductPagination) (products []Product, err error) //[]Product adalah slice untuk mengambil lebih dari satu product kalo cuma Product biasa itu artinya hanya mengembalikan satu Product
	GetProudctBySKU(ctx context.Context, sku string) (product Product, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := NewProductFromCreateProductRequest(req)

	if err = productEntity.Validate(); err != nil {
		log.Log.Errorf(ctx, "[CreateProduct, Validate] with error detail %v", err.Error())
		return
	}

	if err = s.repo.CreateProduct(ctx, productEntity); err != nil {
		return
	}

	return
}

func (s service) ListProducts(ctx context.Context, req ListProductRequestPayload) (products []Product, err error) {
	pagination := NewProductPaginationFromListProductRequest(req)

	products, err = s.repo.GetAllProductWithPaginationCursor(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []Product{}, nil
		}
		return
	}
	if len(products) == 0 {
		return []Product{}, nil
	}
	return
}

func (s service) ProductDetail(ctx context.Context, sku string) (model Product, err error) {
	model, err = s.repo.GetProudctBySKU(ctx, sku)
	if err != nil {
		return
	}
	return
}
