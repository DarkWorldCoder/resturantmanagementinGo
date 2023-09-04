package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type OrderItem struct {
	ID            primitive.ObjectID `bson:"_id"`
	Order_id      string             `json:"order_id" validate:"required"`
	Quantity      *int               `json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Order_item_id string             `json:"order_item_id"`
	Unit_price    *float64           `json:"unit_price" validate:"required"`
	Food_id       *string            `json:"food_id" validate:"required"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
