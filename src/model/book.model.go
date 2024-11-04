package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Display type
type BookDisplay struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

// Default type
type Book struct {
	Id     primitive.ObjectID `json:"id"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
	Year   int                `json:"year"`
	Stock  int                `json:"stock"`
	Price  float32            `json:"price"`
}

// Display all books
type BookDisplayResponse struct {
	DataDisplay []*BookDisplay `json:"datadisplay"`
}

// Display the details of a book
type BookDetailRequest struct {
	Id primitive.ObjectID `json:"id"`
}
type BookDetailResponse struct {
	Data *Book `json:"data"`
}

// Change price and stock values
type BookChangeStockPriceRequest struct {
	Id    primitive.ObjectID `json:"id"`
	Stock int                `json:"stock"`
	Price float32            `json:"price"`
}
type BookChangeStockPriceResponse struct {
	Data *Book `json:"data"`
}

// Add a book
type BookAddRequest struct {
	Data *Book `json:"data"`
}
type BookAddResponse struct {
	Data *Book `json:"data"`
}

// Delete a book
type BookDeleteRequest struct {
	Id primitive.ObjectID `json:"id"`
}
type BookDeleteResponse struct {
	Data *Book `json:"data"`
}
