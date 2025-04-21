package main

import (
	"log"
	"onlineShop/apps/auth"
	"onlineShop/apps/product"
	"onlineShop/apps/transaction"
	"onlineShop/external/database"
	infrafiber "onlineShop/infra/fiber"
	"onlineShop/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	filename := "cmd/api/config.yaml"
	if err := config.LoadConfig(filename); err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	if db != nil {
		log.Println("db connected")
	}

	router := fiber.New(fiber.Config{
		Prefork: true,
		AppName: config.Cfg.App.Name,
	})


	router.Use(infrafiber.Trace())

	auth.Init(router, db) 
	product.Init(router, db)
	transaction.Init(router, db)

	router.Listen(config.Cfg.App.Port)
}
