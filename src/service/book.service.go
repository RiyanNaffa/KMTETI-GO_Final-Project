package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"

	"book-store/src/db"
	"book-store/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Displays all books in the database
func BookDisplayAll() ([]*model.BookDisplay, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")
	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Default().Println("Internal Server Error: Cursor error.")
		return nil, errors.New("internal server error")
	}
	defer cur.Close(context.TODO())

	var bookList []*model.BookDisplay

	for cur.Next(context.TODO()) {
		var book model.Book
		if err := cur.Decode(&book); err != nil {
			log.Default().Println("Internal Server Error: Decoding error.")
			return nil, errors.New("internal server error")
		}
		bookList = append(bookList, &model.BookDisplay{
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		})
	}

	return bookList, nil
}

// Displays a certain book in the database
func BookDetails(idReq *string) (*model.BookDetailed, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")

	var bookDetails *model.Book

	id, err := primitive.ObjectIDFromHex(*idReq)
	if err != nil {
		log.Default().Println("Internal Server Error: Parsing error.")
		return nil, errors.New("internal server error")
	}

	errFind := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&bookDetails)

	if errFind != nil {
		if errors.Is(errFind, mongo.ErrNoDocuments) {
			log.Default().Println("Not Found: Document not found.")
			return nil, errors.New("not found")
		}
		log.Default().Println("Internal Server Error: Decoding error.")
		return nil, errors.New("internal server error")
	}

	return &model.BookDetailed{
		Title:  bookDetails.Title,
		Author: bookDetails.Author,
		Year:   bookDetails.Year,
		Stock:  bookDetails.Stock,
		Price:  bookDetails.Price,
	}, nil
}

// Change a certain price and stock of a book in the database
func BookUpdate(req io.Reader) (*mongo.UpdateResult, error) {
	var changeReq model.BookUpdateRequest

	err := json.NewDecoder(req).Decode(&changeReq)
	if err != nil {
		log.Default().Println("Bad Request: Request body cannot be decoded.")
		return nil, errors.New("bad request")
	}
	id, err := primitive.ObjectIDFromHex(changeReq.Id)
	if err != nil {
		log.Default().Println("Unprocessable Entity: Parsing error.")
		return nil, errors.New("unprocessable entity")
	}
	price, err := primitive.ParseDecimal128(changeReq.Price)
	if err != nil {
		log.Default().Println("Unprocessable Entity: Parsing error.")
		return nil, errors.New("unprocessable entity")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	update := bson.M{
		"$set": bson.M{
			"price": price,
			"stock": changeReq.Stock,
		},
	}

	col := db.MongoDB.Collection("book")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := col.UpdateOne(c, bson.M{"_id": id}, update)
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot update item in database.")
		return nil, errors.New("internal server error")
	}
	return response, nil
}

// Add a book to the database
func BookAdd(req io.Reader) (*mongo.InsertOneResult, error) {
	var bookReq model.BookAddRequest

	err := json.NewDecoder(req).Decode(&bookReq)
	if err != nil {
		log.Default().Println("Bad Request: Request body cannot be decoded.")
		return nil, errors.New("bad request")
	}

	price, err := primitive.ParseDecimal128(bookReq.Price)
	if err != nil {
		log.Default().Println("Unprocessable Entity: Parsing error.")
		return nil, errors.New("unprocessable entity")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := col.InsertOne(c, model.Book{
		Id:     primitive.NewObjectID(),
		Title:  bookReq.Title,
		Author: bookReq.Author,
		Year:   bookReq.Year,
		Stock:  bookReq.Stock,
		Price:  price,
	})
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot add item to database.")
		return nil, errors.New("internal server error")
	}

	return response, nil
}

// Delete a certain book in the database
func BookDelete(idReq *string) (*model.BookDeleteResponse, error) {
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

	var book model.Book

	col := db.MongoDB.Collection("book")

	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	errDel := col.FindOneAndDelete(c, filter).Decode(&book)
	if errDel != nil {
		if errors.Is(errDel, mongo.ErrNoDocuments) {
			log.Default().Println("Not Found: Document not found.")
			return nil, errors.New("not found")
		}
		log.Default().Println("Internal Server Error: Cannot delete the document.")
		return nil, errors.New("internal server error")
	}

	return &model.BookDeleteResponse{
		Id:    *idReq,
		Title: book.Title,
	}, nil
}
