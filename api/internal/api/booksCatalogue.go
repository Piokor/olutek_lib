package api

import (
	"net/http"

	"github.com/Piokor/olutek_lib/internal/bookcatalogue"
	"github.com/Piokor/olutek_lib/internal/googleapi"
	"github.com/gin-gonic/gin"
)

func getBookHandler(c *gin.Context) {
	dbService := getDbService(c)
	bookApiId := c.Param("apiID")

	book, err := bookcatalogue.GetBook(bookApiId, dbService)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBooksQuery(c *gin.Context) *googleapi.BookQuery {
	queryStr := c.Param("query")
	query := googleapi.BookQuery{
		Query:     queryStr,
		Title:     c.Query("title"),
		Author:    c.Query("author"),
		Publisher: c.Query("publisher"),
		Isbn:      c.Query("isbn"),
	}
	return &query
}

func getBooksHandler(c *gin.Context) {
	dbService := getDbService(c)
	booksQuery := getBooksQuery(c)

	books, err := bookcatalogue.GetBooks(booksQuery, dbService)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusOK, books)
}

func bindBooksCatalogueRoutes(authorized gin.IRoutes) {
	authorized.GET("/api/books-catalogue/book/:apiID", getBookHandler)
	authorized.GET("/api/books-catalogue/books/:query", getBooksHandler)
}
