package api

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/gin-gonic/gin"
)

func GetRouter(dbService *database.DbService) *gin.Engine {
	router := gin.Default()
	router.Use(dbServiceMiddleware(dbService))
	router.Use(CorsMiddleware())

	bindAuthRoutes(router)
	authorized := router.Group("/")
	//authorized.Use(authorizedMiddleware())
	bindBooksCatalogueRoutes(authorized)
	bindBooksLibraryRoutes(authorized)

	return router
}
