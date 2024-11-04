package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	Id             primitive.ObjectID `json:"id"`
	Name           string             `json:"name"`
	NIK            string             `json:"nik"`
	Edu            string             `json:"edu"`
	EmploymentDate string             `json:"date"`
	EmploymentType string             `json:"type"`
}

type EmployeeDisplay struct {
	Name           string `json:"name"`
	EmploymentDate string `json:"date"`
	EmploymentType string `json:"type"`
}

type EmployeeDisplayResponse struct {
	DataDisplay []*EmployeeDisplay `json:"datadisplay"`
}

type EmployeeAddRequest struct {
	Data *Employee `json:"data"`
}
type EmployeeAddResponse struct {
	Data *Employee `json:"data"`
}
