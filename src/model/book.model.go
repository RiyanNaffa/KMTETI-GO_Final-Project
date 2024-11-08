package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Database type
type Book struct {
	Id     primitive.ObjectID   `bson:"_id,emitonempty"`
	Title  string               `bson:"title"`
	Author string               `bson:"author"`
	Year   int                  `bson:"year"`
	Stock  int                  `bson:"stock"`
	Price  primitive.Decimal128 `bson:"price"`
}

// Display type
type BookDisplay struct {
	Title  string               `json:"title"`
	Author string               `json:"author"`
	Price  primitive.Decimal128 `json:"price"`
}

// Detailed type
type BookDetailed struct {
	Title  string               `json:"title"`
	Author string               `json:"author"`
	Year   int                  `json:"year"`
	Stock  int                  `json:"stock"`
	Price  primitive.Decimal128 `json:"price"`
}

// Change price and stock values
type BookChangeStockPriceRequest struct {
	Id    primitive.ObjectID   `json:"_id"`
	Stock int                  `json:"stock"`
	Price primitive.Decimal128 `json:"price"`
}
type BookChangeStockPriceResponse struct {
	Data *BookDetailed `json:"data"`
}

// Add a book
type BookAddRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Stock  int    `json:"stock"`
	Price  string `json:"price"`
}

// Delete a book
type BookDeleteRequest struct {
	Id primitive.ObjectID `json:"id"`
}
type BookDeleteResponse struct {
	Data *Book `json:"data"`
}
