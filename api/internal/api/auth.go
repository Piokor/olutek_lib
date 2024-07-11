package api

import (
	"net/http"
	"os"

	"github.com/Piokor/olutek_lib/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type loginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func setCookieAndGoHome(sessionId uuid.UUID, c *gin.Context) {
	c.SetCookie("session", sessionId.String(), 3600, "/", os.Getenv("DOMAIN"), false, false)
	c.Redirect(http.StatusFound, "/")
}

func getRegisterHandler(c *gin.Context) {
	dbService := getDbService(c)

	creds := loginCredentials{}
	err := c.BindJSON(&creds)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	_, sessionId, err := users.RegisterUser(creds.Username, creds.Password, dbService)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	setCookieAndGoHome(*sessionId, c)
}

func getLoginHandler(c *gin.Context) {
	dbService := getDbService(c)

	creds := loginCredentials{}
	err := c.BindJSON(&creds)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	_, sessionId, err := users.LoginUser(creds.Username, creds.Password, dbService)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	setCookieAndGoHome(*sessionId, c)
}

func bindAuthRoutes(router *gin.Engine) {
	router.POST("/api/auth/login", getLoginHandler)
	router.POST("/api/auth/register", getRegisterHandler)
}
