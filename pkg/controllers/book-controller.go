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