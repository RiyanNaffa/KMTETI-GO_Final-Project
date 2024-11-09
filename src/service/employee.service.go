package service

import (
	"book-store/src/db"
	"book-store/src/model"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func EmployeeDisplayAll() ([]*model.EmployeeDisplay, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("employee")
	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Default().Println("Internal Server Error: Cursor error.")
		return nil, errors.New("internal server error")
	}
	defer cur.Close(context.TODO())

	var emplList []*model.EmployeeDisplay

	for cur.Next(context.TODO()) {
		var empl model.Employee
		if err := cur.Decode(&empl); err != nil {
			log.Default().Println("Internal Server Error: Decoding error.")
			return nil, errors.New("internal server error")
		}
		emplList = append(emplList, &model.EmployeeDisplay{
			Name:       empl.Name,
			EmplDate:   empl.EmplDate,
			EmplStatus: empl.EmplStatus,
		})
	}

	return emplList, nil
}

func EmployeeAdd(req io.Reader) (*mongo.InsertOneResult, error) {
	var emplReq model.EmployeeAddRequest

	err := json.NewDecoder(req).Decode(&emplReq)
	if err != nil {
		log.Default().Println("Bad Request: Request body cannot be decoded.")
		return nil, errors.New("bad request")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("employee")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := col.InsertOne(c, model.Employee{
		Id:         primitive.NewObjectID(),
		Name:       emplReq.Name,
		NIK:        emplReq.NIK,
		Edu:        emplReq.Edu,
		EmplDate:   model.Date{Day: emplReq.Day, Month: emplReq.Month, Year: emplReq.Year},
		EmplStatus: emplReq.EmplStatus,
	})
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot add document to database.")
		return nil, errors.New("internal server error")
	}

	return response, nil
}

func EmployeeDelete(idReq *string) (*model.EmployeeDeleteResponse, error) {
	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	id, err := primitive.ObjectIDFromHex(*idReq)
	if err != nil {
		log.Default().Println("Internal Server Error: Parsing Error.")
		return nil, errors.New("internal server error")
	}

	var empl model.Employee

	col := db.MongoDB.Collection("employee")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	errDel := col.FindOneAndDelete(c, filter).Decode(&empl)
	if errDel != nil {
		if errors.Is(errDel, mongo.ErrNoDocuments) {
			log.Default().Println("Not Found: Document not found.")
			return nil, errors.New("not found")
		}
		log.Default().Println("Internal Server Error: Cannot delete the document.")
		return nil, errors.New("internal server error")
	}

	return &model.EmployeeDeleteResponse{
		Id:   *idReq,
		Name: empl.Name,
	}, nil
}
