package router

import (
	"TMD-Backend-Go/controller"
	"TMD-Backend-Go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUpRouter(app *gin.Engine, db *mongo.Database) *gin.Engine {

	mainRouter := app

	userController := controller.NewController(db, models.User{})
	todoController := controller.NewController(db, models.Todo{})

	mainRouter.POST("/login", userController.LoginHandler)
	mainRouter.POST("/register", userController.RegisterHandler)

	userRouter := mainRouter.Group("/user")
	{

		userRouter.GET("/", userController.GetUsersHandler)
		userRouter.GET("/:userid", userController.GetUserHandler)

		userRouter.POST("/", userController.PostUsersHandler)
		userRouter.POST("/:userid", userController.PostUserHandler)
	}

	todoRouter := mainRouter.Group("/todo")
	{
		todoRouter.GET("/", todoController.GetTodosHandler)
		todoRouter.GET("/:todoid", todoController.GetTodoHandler)

		todoRouter.POST("/", todoController.PostTodosHandler)
		todoRouter.POST("/:todoid", todoController.PostTodoHandler)
	}

	return mainRouter
}
