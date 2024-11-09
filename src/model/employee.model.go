package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Date struct {
	Day   int `json:"day" bson:"day"`
	Month int `json:"month" bson:"month"`
	Year  int `json:"year" bson:"year"`
}

// Type used for decoding a database document.
type Employee struct {
	Id         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	NIK        string             `bson:"nik"`
	Edu        string             `bson:"edu"`
	EmplDate   Date               `bson:"date"`
	EmplStatus string             `bson:"type"`
}

// Type used for displaying general information regarding an employee document.
type EmployeeDisplay struct {
	Name       string `json:"name"`
	EmplDate   Date   `json:"date"`
	EmplStatus string `json:"type"`
}

// Type used for decoding a create request of an employee document.
type EmployeeAddRequest struct {
	Name       string `json:"name"`
	NIK        string `json:"nik"`
	Edu        string `json:"edu"`
	Day        int    `json:"day"`
	Month      int    `json:"month"`
	Year       int    `json:"year"`
	EmplStatus string `json:"type"`
}

// Type used for a response to delete an employee document procedure.
type EmployeeDeleteResponse struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}
