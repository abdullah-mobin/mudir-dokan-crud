package handler

import (
	"fmt"
	"log"
	"mudir-dokan-crud/data"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type DataStruct struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}

func GetAllItems(c *fiber.Ctx) error {

	item := data.GetAllItems()
	response := Response{
		Status:  true,
		Code:    200,
		Message: "All item found",
		Data:    item,
	}
	return c.JSON(response)
}

func CreateItem(c *fiber.Ctx) error {

	var insertData DataStruct
	if err := c.BodyParser(&insertData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid post req")
	}

	data.CreateItem(insertData.Name, insertData.Quantity, insertData.Price)

	return c.JSON(fiber.NewError(200, "item created succesfully"))
}

func UpdateItemById(c *fiber.Ctx) error {

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	var i DataStruct
	if err := c.BodyParser(&i); err != nil {
		log.Fatalf("format %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "invalid update req")
	}

	data.UpdateItemById(id, i.Name, i.Quantity, i.Price)

	return c.JSON((fiber.NewError(200, "Item updated succesfully")))
}

func DeleteItemById(c *fiber.Ctx) error {

	param := c.Params("id")
	id, _ := strconv.Atoi(param)

	err := data.DeleteItemById(id)
	if err != nil {
		return fmt.Errorf("error deleting items: %v", err)
	}
	return c.JSON(fiber.NewError(200, "item deleted succesfully"))
}
