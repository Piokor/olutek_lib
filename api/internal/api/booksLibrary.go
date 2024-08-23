package api

import (
	"net/http"

	"github.com/Piokor/olutek_lib/internal/library"
	"github.com/gin-gonic/gin"
)

type addToLibraryData struct {
	CatalogueBookId string `json:"catalogue_book_id"`
}

func addBookToLibrary(c *gin.Context) {
	dbService := getDbService(c)
	userId := c.MustGet("userId").(int64)

	var addData addToLibraryData
	err := c.BindJSON(&addData)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	err = library.AddBookToLibrary(addData.CatalogueBookId, userId, dbService)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.String(http.StatusOK, "Book added successfuly")
}

func bindBooksLibraryRoutes(authorized gin.IRoutes) {
	authorized.POST("/api/books-library/add-book", addBookToLibrary)
}
