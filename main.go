package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// repsentasi data
type Book struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// slice di go
var books = []Book{
	{ID: "1", Name: "akramulfata", Description: "buku keren", Price: 5.00},
	{ID: "2", Name: "fata", Description: "buku jelek", Price: 8.00},
	{ID: "3", Name: "progremming", Description: "belajar koding", Price: 76.99},
}

// getBooks responds with the list of all albums as json
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// add buku baru
func CreateBook(c *gin.Context) {
	var addBook Book

	if err := c.BindJSON(&addBook); err != nil {
		return
	}

	books = append(books, addBook)
	c.IndentedJSON(http.StatusCreated, addBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", CreateBook)
	router.GET("/books/:id", GetBookByID)

	router.Run("localhost:8080")
}
