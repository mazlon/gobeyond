package router

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type Cookies map[string]interface{}

type GbtServer struct {
	dbConnection *sql.DB
	router       *gin.Engine
}

func NewGbtServer(dbConnection *sql.DB, router *gin.Engine) *GbtServer {
	apiServer := GbtServer{dbConnection: dbConnection, router: router}
	apiServer.PrepareRoutest()
	return &apiServer

}

func SetCookie(c *gin.Context, cookies Cookies) error {
	for key, value := range cookies {
		value, ok := value.(string)
		if !ok {
			return errors.New("error while setting cookie")

		}
		c.SetCookie(key, value, 1000000, "/", "", true, true)
	}
	return nil
}


