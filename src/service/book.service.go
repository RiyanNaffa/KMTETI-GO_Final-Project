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
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Displays all books in the database
func BookDisplayAll() ([]*model.BookDisplay, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error: cannot connect to database")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")
	projection := bson.D{
		{Key: "_id", Value: 0},
		{Key: "title", Value: 1},
		{Key: "author", Value: 1},
		{Key: "price", Value: 1},
	}
	cur, err := col.Find(context.TODO(), bson.D{}, options.Find().SetProjection(projection))

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error: cursor error")
	}
	defer cur.Close(context.TODO())

	var bookList []*model.BookDisplay

	for cur.Next(context.TODO()) {
		var book model.Book
		if err := cur.Decode(&book); err != nil {
			log.Default().Println(err.Error())
			return nil, errors.New("internal server error: decoding error")
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
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error: cannot connect to database")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")

	var bookDetails *model.Book
	id, err := primitive.ObjectIDFromHex(*idReq)
	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error: cannot parse string ID to ObjectID")
	}

	errFind := col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&bookDetails)

	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			log.Default().Println(errFind.Error())
			return nil, errors.New("document not found")
		}
		log.Default().Println(errFind.Error())
		return nil, errors.New("internal server error: decoding error")
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
// func ChangePrice(req io.Reader) (*model.BookChangeStockPriceResponse, error)

// Add a book to the database
func BookAdd(req io.Reader) (*mongo.InsertOneResult, error) {
	var bookReq model.BookAddRequest

	err := json.NewDecoder(req).Decode(&bookReq)
	if err != nil {
		log.Default().Println("Bad Request: Request body cannot be decoded.")
		return nil, errors.New("bad request")
	}
	decimalValue, err := primitive.ParseDecimal128(bookReq.Price)
	if err != nil {
		log.Default().Println("Internal Server Error: Parsing error.")
		return nil, errors.New("parsing error")
	}

	db, err := db.DBConnection()
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot connect to database.")
		return nil, errors.New("internal server error: cannot connect to database")
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
		Price:  decimalValue,
	})
	if err != nil {
		log.Default().Println("Internal Server Error: Cannot add item to database.")
		return nil, errors.New("internal server error")
	}

	return response, nil
}

// Delete a certain book in the database
// func DeleteBook(req io.Reader) (*model.BookDeleteResponse, error){
// 	var bookDelReq model.BookDeleteRequest
// 	errDec := json.NewDecoder(req).Decode(&bookDelReq)

// 	if errDec != nil{
// 		return nil,
// 	}
// }
