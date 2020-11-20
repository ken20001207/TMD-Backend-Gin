package controller

import (
	"TMD-Backend-Go/models"
	"TMD-Backend-Go/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type Response struct {
	Message string `json:"message"`
}

func (r *Controller) GetUsersHandler(c *gin.Context) {

	userCollection := r.DB.Collection("user")
	ctx := context.Background()
	cur, err := userCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var res []models.User

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result models.User
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, res)
}

func (r *Controller) PostUsersHandler(c *gin.Context) {

}

func (r *Controller) GetUserHandler(c *gin.Context) {
	res := models.User{
		PasswordMD5: "passwordMd5",
		NickName:    "nickname",
	}

	c.JSON(200, res)

}

func (r *Controller) PostUserHandler(c *gin.Context) {

}

func (r *Controller) LoginHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	ctx := context.Background()
	var user models.User
	res := r.DB.Collection("user").FindOne(ctx, bson.D{{"email", email}})

	err := res.Decode(&user)

	if err != nil {
		_ = c.Error(fmt.Errorf("EMAIL_OR_PASSWORD_WRONG"))
		return
	}

	if !utils.CheckPasswordHash(password, user.PasswordMD5) {
		_ = c.Error(fmt.Errorf("EMAIL_OR_PASSWORD_WRONG"))
		return
	}

	c.JSON(200, user)
}

func (r *Controller) RegisterHandler(c *gin.Context) {
	userCollection := r.DB.Collection("user")

	ctx := context.Background()

	passwordMd5, err := utils.HashPassword(c.PostForm("password"))

	one, err := userCollection.InsertOne(ctx, models.User{
		Email:       c.PostForm("email"),
		NickName:    c.PostForm("nickname"),
		PasswordMD5: passwordMd5,
	})

	if err != nil {
		log.Fatal(err)
	}

	var createdUser models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": one.InsertedID}).Decode(&createdUser)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, createdUser)
}
