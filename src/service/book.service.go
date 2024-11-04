package service

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"

	"book-store/src/db"
	"book-store/src/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Displays all books in the database
func DisplayAllBooks() (*model.BookDisplayResponse, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("bookdisplay")
	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var bookList []*model.BookDisplay

	for cur.Next(context.TODO()) {
		var book model.BookDisplay
		cur.Decode(&book)
		bookList = append(bookList, &model.BookDisplay{
			Title:  book.Title,
			Author: book.Title,
			Price:  book.Price,
		})
	}

	return &model.BookDisplayResponse{
		DataDisplay: bookList,
	}, nil
}

// Displays a certain book in the database
func DetailDisplay() (*model.BookDetailResponse, error) {
	db, errCon := db.DBConnection()

	if errCon != nil {
		log.Default().Println(errCon.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")
	cur, err := col.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}

	var bookDetails *model.Book

	cur.Next(context.TODO())
	var book model.Book
	cur.Decode(&book)
	bookDetails = &model.Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
		Year:   book.Year,
		Stock:  book.Stock,
		Price:  book.Price,
	}

	return &model.BookDetailResponse{
		Data: bookDetails,
	}, nil
}

// Change a certain price and stock of a book in the database
func ChangePrice(req io.Reader) (*model.BookChangeStockPriceResponse, error)

// Add a book to the database
func AddBook(req io.Reader) (*model.BookAddResponse, error) {
	var bookReq model.BookAddRequest
	errReq := json.NewDecoder(req).Decode(&bookReq)
	if errReq != nil {
		return nil, errors.New("bad request")
	}

	db, errCon := db.DBConnection()
	if errCon != nil {
		log.Default().Println(errCon.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")
	colDisp := db.MongoDB.Collection("bookdisplay")

	_, errIns := col.InsertOne(context.TODO(), model.Book{
		Id:     primitive.NewObjectID(),
		Title:  bookReq.Data.Title,
		Author: bookReq.Data.Author,
		Year:   bookReq.Data.Year,
		Stock:  bookReq.Data.Stock,
		Price:  bookReq.Data.Price,
	})
	_, errInsDisp := colDisp.InsertOne(context.TODO(), model.BookDisplay{
		Title:  bookReq.Data.Title,
		Author: bookReq.Data.Author,
		Price:  bookReq.Data.Price,
	})
	if errIns != nil {
		log.Default().Println(errIns.Error())
		return nil, errors.New("internal server error")
	}
	if errInsDisp != nil {
		log.Default().Println(errInsDisp.Error())
		return nil, errors.New("internal server error")
	}

	return nil, nil
}

// Delete a certain book in the database
func DeleteBook(req io.Reader) (*model.BookDeleteResponse, error)
