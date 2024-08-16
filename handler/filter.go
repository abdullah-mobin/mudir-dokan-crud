package handler

import (
	"database/sql"
	"fmt"
	"mudir-dokan-crud/data"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool        `josn:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `josn:"data"`
}

func WelcomeMsg(c *fiber.Ctx) error {
	return fmt.Errorf("welcome to new app")
}

func GetItemById(c *fiber.Ctx) error {
	param := c.Params("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadRequest, "invalid id formate"))
	}

	item, err := data.GetItemById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response := Response{
				Status:  false,
				Code:    fiber.StatusNotFound,
				Message: "Item not found",
				Data:    item,
			}
			return c.JSON(response)
		}
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error"))

	}
	response := Response{
		Status:  true,
		Code:    200,
		Message: "Item found",
		Data:    item,
	}

	return c.JSON(response)
}

func GetQuantityById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadRequest, "invalid id formate"))
	}

	quantity, err := data.GetCurrentQuantityById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response := Response{
				Status:  false,
				Code:    fiber.StatusNotFound,
				Message: "quantity not found",
				Data:    quantity,
			}
			return c.JSON(response)
		}
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, "internal server error"))
	}
	response := Response{
		Status:  true,
		Code:    200,
		Message: "Quantity found",
		Data:    quantity,
	}

	return c.JSON(response)
}

func GetPriceById(c *fiber.Ctx) error {
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusBadRequest, "invalid id formate"))

	}

	price, err := data.GetCurrentPriceById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			response := Response{
				Status:  false,
				Code:    fiber.StatusNotFound,
				Message: "Price not found",
				Data:    price,
			}
			return c.JSON(response)
		}
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, "internal server error"))
	}

	response := Response{
		Status:  true,
		Code:    fiber.StatusOK,
		Message: "Price found",
		Data:    price,
	}

	return c.JSON(response)
}
