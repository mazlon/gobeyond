package router

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/vgarvardt/gue/v5"
)

type Cookies map[string]interface{}

type GbtServer struct {
	dbConnection *sql.DB
	router       *gin.Engine
	queue        *gue.Client
}

func NewGbtServer(dbConnection *sql.DB, router *gin.Engine, gue *gue.Client) *GbtServer {
	apiServer := GbtServer{dbConnection: dbConnection, router: router, queue: gue}
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
