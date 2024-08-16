package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Items struct {
	Id       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
	Price    float32 `json:"price,omitempty"`
	Total    float32 `json:"total,omitempty"`
}

var DB *sql.DB

func CreateItem(name string, quantity int, price float32) {

	insertQuery := `INSERT INTO items (name, quantity, price, total) VALUES (?, ?, ?, ?)`

	DB.Exec(insertQuery, name, quantity, price, float32(quantity)*price)
	ReindexItems()
	fmt.Printf("Data inserted successfully.")
}

func GetAllItems() []Items {

	fmt.Println("connected successfully.")

	readQuery := `SELECT id, name, quantity, price, total FROM items`

	res, err := DB.Query(readQuery)

	var item []Items
	fmt.Println("query successfully.")

	for res.Next() {
		var i Items
		if err := res.Scan(&i.Id, &i.Name, &i.Quantity, &i.Price, &i.Total); err != nil {
			fmt.Println("row contains error")
		}
		item = append(item, i)
	}

	if err != nil {
		fmt.Println("error query")
	}

	fmt.Println("read successful", item)
	return item
}

func UpdateItemById(id int, name string, quantity int, price float32) error {
	queryClauses := []string{}
	args := []interface{}{}

	if name != "" {
		queryClauses = append(queryClauses, "name = ?")
		args = append(args, name)
	}

	if quantity != 0 {
		queryClauses = append(queryClauses, "quantity = ?")
		args = append(args, quantity)
	}
	if price != 0 {
		queryClauses = append(queryClauses, "price = ?")
		args = append(args, price)
	}
	if len(queryClauses) == 0 {
		return fmt.Errorf("no field to update")
	}

	var total float32
	total = price * float32(quantity)
	if price == 0 {
		p, _ := GetCurrentPriceById(id)
		total = p.Price * float32(quantity)
	}
	if quantity == 0 {
		q, _ := GetCurrentQuantityById(id)
		total = float32(q.Quantity) * price
	}

	if price == 0 && quantity == 0 {
		q, _ := GetCurrentQuantityById(id)
		p, _ := GetCurrentPriceById(id)
		total = p.Price * float32(q.Quantity)
	}

	queryClauses = append(queryClauses, "total = ?")
	args = append(args, total)

	query := fmt.Sprintf("UPDATE items SET %s WHERE id = ?", strings.Join(queryClauses, ", "))
	args = append(args, id)

	_, err := DB.Exec(query, args...)
	if err != nil {
		fmt.Println("error in execution update")
		log.Fatalf("Error opening database: %v", err)
	}

	fmt.Println("t p ", total)

	return fmt.Errorf("updated")
}

func DeleteItemById(id int) error {

	deleteQuery := `DELETE FROM items WHERE id = ?`

	_, err := DB.Exec(deleteQuery, id)

	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}
	ReindexItems()
	fmt.Println("item deleted succesfully")
	return nil
}
func ReindexItems() error {

	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
	}
	_, err = tx.Exec("SET @new_id = 0;")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error setting initial ID: %v", err)
	}
	_, err = tx.Exec("UPDATE items SET id = @new_id := @new_id + 1 ORDER BY id;")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating IDs: %v", err)
	}
	_, err = tx.Exec("ALTER TABLE items AUTO_INCREMENT = 1;")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error resetting auto-increment: %v", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}
	log.Println("Successfully reindexed items.")
	return nil
}
