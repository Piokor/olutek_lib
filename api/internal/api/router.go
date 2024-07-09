package api

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/gin-gonic/gin"
)

func GetRouter(dbService *database.DbService) *gin.Engine {
	router := gin.Default()
	router.Use(dbServiceMiddleware(dbService))

	router.GET("/api/books-catalogue/book/:apiID", getBookHandler)
	router.GET("/api/books-catalogue/books/:query", getBooksHandler)

	return router
}
