package api

import (
	"net/http"

	"github.com/Piokor/olutek_lib/internal/database"
	"github.com/Piokor/olutek_lib/internal/session"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func authorizedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionCookie, err := c.Cookie("session")
		if err != nil {
			c.String(http.StatusUnauthorized, "Session cookie missing")
			c.Abort()
			return
		}
		sessionId, err := uuid.Parse(sessionCookie)
		if err != nil {
			c.String(http.StatusUnauthorized, "Invalid session id")
			c.Abort()
			return
		}
		dbService := getDbService(c)
		userId, err := session.GetUserId(sessionId, dbService)
		if err != nil {
			c.String(http.StatusInternalServerError, "Wrong middlware order")
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}
