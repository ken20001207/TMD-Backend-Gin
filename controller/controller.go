package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type Controller struct {
	DB      *mongo.Database
	New     func() Resource
	Timeout time.Duration
	Limit   int64
}

func NewController(db *mongo.Database, resource interface{}) *Controller {
	return &Controller{
		DB: db,
		New: func() Resource {
			return reflect.New(reflect.TypeOf(resource)).Interface().(Resource)
		},
		Timeout: 60 * time.Second,
		Limit:   1000,
	}
}

type Decoder interface {
	Decode(val interface{}) error
}

type Resource interface {
	SetID(id primitive.ObjectID)
	Valid() error
	Decode(cursor Decoder) error
}
