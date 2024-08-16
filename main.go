package main

import (
	"database/sql"
	"fmt"
	"log"
	"mudir-dokan-crud/data"
	"mudir-dokan-crud/routes"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	schema := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, schema)

	data.DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer data.DB.Close()
	fmt.Println("Welcome to Mudir Dokan CRUD API")

	app := fiber.New()
	routes.SetupRoutes(app)

	app.Listen(":8080")
}
