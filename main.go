package main

import (
	"TMD-Backend-Go/controller"
	"TMD-Backend-Go/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func setUpEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func setUpRouter(db *mongo.Database) *gin.Engine {
	mainRouter := gin.Default()

	userRouter := mainRouter.Group("/user")
	{
		userController := controller.NewController(db, models.User{})

		userRouter.GET("/", userController.GetUsersHandler)
		userRouter.GET("/:userid", userController.GetUserHandler)

		userRouter.POST("/", userController.PostUsersHandler)
		userRouter.POST("/:userid", userController.PostUserHandler)
	}

	todoRouter := mainRouter.Group("/todo")
	{
		todoController := controller.NewController(db, models.Todo{})

		todoRouter.GET("/", todoController.GetTodosHandler)
		todoRouter.GET("/:todoid", todoController.GetTodoHandler)

		todoRouter.POST("/", todoController.PostTodosHandler)
		todoRouter.POST("/:todoid", todoController.PostTodoHandler)
	}

	return mainRouter
}

func setUpMongoClient() *mongo.Database {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	mongoUri := "mongodb+srv://" + dbUser + ":" + dbPass + "@" + dbHost

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbName)
}

func main() {

	setUpEnv()

	mongoClient := setUpMongoClient()

	mainRouter := setUpRouter(mongoClient)

	log.Fatal(mainRouter.Run(":8080"))
}
