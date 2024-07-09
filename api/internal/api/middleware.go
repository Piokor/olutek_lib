package api

import (
	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/gin-gonic/gin"
)

func dbServiceMiddleware(dbService *database.DbService) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("dbService", dbService)
		c.Next()
	}
}

func getDbService(c *gin.Context) *database.DbService {
	return c.MustGet("dbService").(*database.DbService)
}
