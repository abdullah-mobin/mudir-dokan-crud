package data

import (
	_ "github.com/go-sql-driver/mysql"
)

func GetItemById(id int) (Items, error) {

	var item Items

	query := `SELECT id, name, quantity, price, total FROM items WHERE id = ?`
	res := DB.QueryRow(query, id)
	err := res.Scan(&item.Id, &item.Name, &item.Quantity, &item.Price, &item.Total)

	if err != nil {
		return item, err
	}

	return item, nil
}

func GetCurrentQuantityById(id int) (Items, error) {
	var quantity Items

	query := `SELECT quantity FROM items WHERE id = ?`
	res := DB.QueryRow(query, id)
	err := res.Scan(&quantity.Quantity)

	if err != nil {
		return quantity, err
	}

	return quantity, nil
}

func GetCurrentPriceById(id int) (Items, error) {

	var price Items

	query := `SELECT price FROM items WHERE id = ?`
	res := DB.QueryRow(query, id)
	err := res.Scan(&price.Price)
	if err != nil {
		return price, err
	}

	return price, nil
}

func GetCurrentTotalPriceById(id int) float32 {

	var current float32
	query := `SELECT total FROM items WHERE id = ?`
	res := DB.QueryRow(query, id)
	res.Scan(&current)

	return current
}
