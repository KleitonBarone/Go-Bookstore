package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/KleitonBarone/Go-Bookstore/pkg/utils"
	"github.com/KleitonBarone/Go-Bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(response http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	responseBody, _ := json.Marshal(newBooks)
	response.Header().Set("Content-Type","Application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(responseBody)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	responseBody, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type","Application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(responseBody)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	bookData := &models.Book{}
	utils.ParseBody(request, bookData)
	validationErrors := []string{}
	if bookData.Name == "" {
		validationErrors = append(validationErrors, fmt.Errorf("Name is required").Error())
	}
	if bookData.Author == "" {
		validationErrors = append(validationErrors, fmt.Errorf("Author is required").Error())
	}
	if bookData.Publication == "" {
		validationErrors = append(validationErrors, fmt.Errorf("Publication is required").Error())
	}
	if len(validationErrors) > 0 {
		response.Header().Set("Content-Type","Application/json")
		response.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(response).Encode(validationErrors); err != nil {}
		return
	}
	newBook := bookData.CreateBook()
	responseBody, _ := json.Marshal(newBook)
	response.Header().Set("Content-Type","Application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(responseBody)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	deletedBook := models.DeleteBook(Id)
	responseBody, _ := json.Marshal(deletedBook)
	response.Header().Set("Content-Type","Application/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(responseBody)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	bookData := &models.Book{}
	utils.ParseBody(request, bookData)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(Id)
	if bookData.Name != "" {
		bookDetails.Name = bookData.Name
	}
	if bookData.Author != "" {
		bookDetails.Author = bookData.Author
	}
	if bookData.Publication != "" {
		bookDetails.Publication = bookData.Publication
	}
	db.Save(&bookDetails)
	responseBody, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type","Application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(responseBody)
}
