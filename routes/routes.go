package routes

import (
	"mudir-dokan-crud/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/get-all-items", handler.GetAllItems).Name("get-all-items")
	v1.Get("/get-item-by-id/:id", handler.GetItemById).Name("get-item-by-id")
	v1.Get("/get-quantity-by-id/:id", handler.GetQuantityById).Name("get-quantity-by-id")
	v1.Get("/get-price-by-id/:id", handler.GetPriceById).Name("get-price-by-id")

	v1.Post("/create-item", handler.CreateItem).Name("create-item")
	v1.Put("/update-item-by-id/:id", handler.UpdateItemById).Name("update-item-by-id")
	v1.Delete("/delete-item/:id", handler.DeleteItemById).Name("delete-item-by-id")

}
