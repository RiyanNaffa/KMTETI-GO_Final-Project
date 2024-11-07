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
func BookDetails(idStr *string) (*model.BookDetailed, error) {
	db, err := db.DBConnection()

	if err != nil {
		log.Default().Println(err.Error())
		return nil, errors.New("internal server error: cannot connect to database")
	}
	defer db.MongoDB.Client().Disconnect(context.TODO())

	col := db.MongoDB.Collection("book")

	var bookDetails *model.Book
	id, err := primitive.ObjectIDFromHex(*idStr)
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
func BookAdd(req io.Reader) (*model.BookAddResponse, error) {
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
// func DeleteBook(req io.Reader) (*model.BookDeleteResponse, error){
// 	var bookDelReq model.BookDeleteRequest
// 	errDec := json.NewDecoder(req).Decode(&bookDelReq)

// 	if errDec != nil{
// 		return nil,
// 	}
// }
