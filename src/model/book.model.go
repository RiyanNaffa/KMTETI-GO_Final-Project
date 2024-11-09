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

// BookDisplayAll()
//
// Data type for "display" response
type BookDisplay struct {
	Title  string               `json:"title"`
	Author string               `json:"author"`
	Price  primitive.Decimal128 `json:"price"`
}

// BookDetails()
// Data type for "details" response
type BookDetailed struct {
	Title  string               `json:"title"`
	Author string               `json:"author"`
	Year   int                  `json:"year"`
	Stock  int                  `json:"stock"`
	Price  primitive.Decimal128 `json:"price"`
}

// BookUpdate()
// Data type for "change" request
type BookUpdateRequest struct {
	Id    string `json:"_id"`
	Stock int    `json:"stock"`
	Price string `json:"price"`
}

// BookAdd()
// Data type for "add" request
type BookAddRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Stock  int    `json:"stock"`
	Price  string `json:"price"`
}

// BookDelete()
// Data type for "delete" response
type BookDeleteResponse struct {
	Id    string `json:"_id"`
	Title string `json:"title"`
}
