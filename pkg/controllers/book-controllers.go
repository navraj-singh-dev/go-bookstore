package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/navraj-singh-dev/go-bookstore/pkg/models"
	utils "github.com/navraj-singh-dev/go-bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var createdBookStruct *models.BookModel
	utils.UnmarshallBody(r, createdBookStruct)
	createdBookStruct = models.CreateBook(createdBookStruct)

	// marshall
	createdBookByte, _ := json.Marshal(createdBookStruct)

	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(createdBookByte)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// run database query to get all books
	allBooksStruct := models.GetAllBooks()

	// convert to []byte, marshalling
	allBooksByte, _ := json.Marshal(allBooksStruct)

	// write header information to response
	w.Header().Set("Content-Type", "application/json")

	// write status code to response
	w.WriteHeader(http.StatusOK)

	// write the main juicy data in JSON.. back to response
	w.Write(allBooksByte)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// get string id from params
	var id string
	params := mux.Vars(r)
	id = params["id"]

	// convert id to int64
	idInt, _ := strconv.ParseInt(id, 0, 0)

	// get the book struct
	retrievedBookStruct, _ := models.GetBookById(idInt)

	// marshall retrieved data
	retrievedBookByte, _ := json.Marshal(retrievedBookStruct)

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(retrievedBookByte)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	// get string id from params
	params := mux.Vars(r)
	id := params["id"]

	// convert id to int
	idInt, _ := strconv.ParseInt(id, 0, 0)

	// delete book
	deletedBookStruct := models.DeleteById(idInt)

	// marshall
	deletedBookByte, _ := json.Marshal(deletedBookStruct)

	// send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deletedBookByte)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {

	// convert request body to struct (unmarshall)
	var updateDetails *models.BookModel
	_ = utils.UnmarshallBody(r, &updateDetails)

	// get id from req body
	params := mux.Vars(r)
	idString := params["id"]

	// convert id to int
	idInt, _ := strconv.ParseInt(idString, 0, 0)

	// UPDATE THE BOOK
	// get book from database
	bookStruct, db := models.GetBookById(idInt)
	// update details
	if updateDetails.Name != "" {
		bookStruct.Name = updateDetails.Name
	}
	if updateDetails.Publication != "" {
		bookStruct.Publication = updateDetails.Publication
	}
	if updateDetails.Author != "" {
		bookStruct.Author = updateDetails.Author
	}
	// save the updated book to db
	db.Save(&bookStruct)
	// convert to []byte
	updatedBookByte, _ := json.Marshal(bookStruct)

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(updatedBookByte)
}
