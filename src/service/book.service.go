package service

import (
	"context"
	// "encoding/json"
	"errors"
	"io"
	"log"

	"book-store/src/db"
	"book-store/src/model"

	"go.mongodb.org/mongo-driver/bson"
)

func DisplayAllBooks() (*model.BookDisplayResponse, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")
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
func DetailDisplay() (*model.BookDetailResponse, error)
func ChangePrice(req io.Reader) (*model.BookChangeStockPriceResponse, error)
func AddBook(req io.Reader) (*model.BookAddResponse, error)
func DeleteBook(req io.Reader) (*model.BookDeleteResponse, error)
