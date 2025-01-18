package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/SamraAzizi/GoProjects/pkg/models"
	"github.com/akhil/go-bookstore/pkg/models"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
