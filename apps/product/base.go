package product

import (
	infrafiber "onlineShop/infra/fiber"
	"onlineShop/apps/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Init(router fiber.Router, db *sqlx.DB) {
	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	productRoute := router.Group("products")
	{
		productRoute.Get("", handler.GetListProducts)
		productRoute.Get("/sku/:sku", handler.GetProductDetail)

		// need authorization

		productRoute.Post("", infrafiber.CheckAuth(), infrafiber.CheckRoles([]string{string(auth.ROLE_Admin)}), handler.CreateProduct)
	}
}
