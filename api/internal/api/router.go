package api

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/gin-gonic/gin"
)

func GetRouter(dbService *database.DbService) *gin.Engine {
	router := gin.Default()
	router.Use(dbServiceMiddleware(dbService))

	bindAuthRoutes(router)
	authorized := router.Group("/")
	authorized.Use(authorizedMiddleware())
	bindBooksCatalogueRoutes(authorized)

	return router
}
