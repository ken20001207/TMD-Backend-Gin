package main

import (
	. "TMD-Backend-Go/db"
	. "TMD-Backend-Go/middleware"
	. "TMD-Backend-Go/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func setUpEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	app := gin.Default()

	app.Use(ErrorHandler())

	setUpEnv()
	mongoClient := SetUpMongoClient()
	SetUpRouter(app, mongoClient)

	log.Fatal(app.Run(":8080"))
}
