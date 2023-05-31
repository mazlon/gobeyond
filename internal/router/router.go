package router

import (
	"github.com/mazlon/gobeyond/internal/handlers"
	"github.com/gin-gonic/gin"
	"database/sql"
)

func PrepareRoutes(r *gin.Engine, db sql.DB) {
	r.GET("/hello", handlers.HelloWorld)
	r.POST("/ask", handlers.AskQuestions(db))
}
