package main

import (
	"asofianidis/collaboraTask/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"log"
)

var DB *gorm.DB

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(logger.New())
	app.Use(recover.New())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Err loading .env file")
	}

	db, dbErr := database.ConnectToDatabase()
	database.MigrateModels(db)

	if dbErr != nil {
		log.Fatal("Error connecting to database")
	}

	DB = db
	fmt.Println("Connected to database")

	app.Listen(":8080")
}
